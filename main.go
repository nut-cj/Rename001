package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Ext will return fileName, extensionName (xxx)(.xxx)
func Ext(path string) (string, string) {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[:i], path[i:]
		}
	}
	return "", ""
}

func main() {
	characterAmount, _ := strconv.Atoi(os.Args[1])
	dir, error := os.OpenFile(".", os.O_RDONLY, os.ModeDir)
	if error != nil {
		defer dir.Close()
		fmt.Println(error.Error())
		return
	}
	names, _ := dir.Readdir(-1)
	for _, name := range names {
		fileName, extensionName := Ext(name.Name())
		reg := regexp.MustCompile("[0-9]+$")
		res := reg.FindAllString(fileName, -1)
		numberName := "00000" + res[len(res)-1]
		var substitution = numberName[len(numberName)-characterAmount:]
		fileName = reg.ReplaceAllString(fileName, substitution)
		finalName := fileName + extensionName
		fmt.Println(finalName, "<--", name.Name())
		os.Rename("./"+name.Name(), "./"+finalName)

	}
}
