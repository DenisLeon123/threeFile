package main

import (
	"fmt"
	"io"
	"os"
	//"path/filepath"
	"io/ioutil"
	//"strings"
	"strconv"
)

func main() {
	out := os.Stdout
	//if !(len(os.Args) == 2 || len(os.Args) == 3) {
	//	panic("usage go run main.go . [-f]")
	//}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	var sizeString string
	var quantity int
	ots := "|	"
	delS := false
	number := 1
	pointer := "├───"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	if printFiles == false {
		for _, file := range files {
			if file.IsDir() {
				quantity++
			}
		}
	} else {
		quantity = len(files)
	}

	for _, file := range files {

		if (printFiles != true && file.IsDir() == true) || (printFiles == true) {

			if number == quantity {
				pointer = "└───"
				delS = true
				if printFiles == false {
					ots = "		"
				} else {
					ots = "	"
				}
			}

			if file.IsDir() == true {
				fmt.Println(pointer + file.Name())
			} else {

				size := file.Size()

				if size == 0 {
					sizeString = "empty"
				} else {
					sizeString = strconv.FormatInt(size, 10) + "b"
				}
				fmt.Println(pointer + file.Name() + " (" + sizeString + ")")
			}

			if file.IsDir() == true {
				pathFolder := path + "/" + file.Name()
				imputThree(file, pathFolder, ots, printFiles, delS)
			}
		}
		number++
	}
	return nil
}

func imputThree(file os.FileInfo, pathFolder string, Three string, printFiles bool, delS bool) error {

	var sizeString string
	ots := "|	"
	number := 1
	pointer := "├───"

	if delS {
		Three = Three[1:]
	}

	files, err := ioutil.ReadDir(pathFolder)
	if err != nil {
		fmt.Println(err)
	}

	quantity := len(files)

	for _, file := range files {

		if (printFiles != true && file.IsDir() == true) || (printFiles == true) {

			if number == quantity {
				pointer = "└───"

				if printFiles == false {
					ots = "		"
				} else {
					ots = "	"
				}
			}

			if file.IsDir() == true {
				fmt.Println(Three + pointer + file.Name())
			} else {

				size := file.Size()

				if size == 0 {
					sizeString = "empty"
				} else {
					sizeString = strconv.FormatInt(size, 10) + "b"
				}

				fmt.Println(Three + pointer + file.Name() + " (" + sizeString + ")")
			}

			if file.IsDir() == true {
				pathFolder2 := pathFolder + "/" + file.Name()
				imputThree(file, pathFolder2, Three+ots, printFiles, delS)
			}
		}
		number++
	}
	return nil
}
