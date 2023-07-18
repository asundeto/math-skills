package main

import (
	"fmt"
	maski "maski/functions"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error! Incorect parametrs count!")
		return
	}
	if CheckTxt(os.Args[1]) {
		fmt.Println("Error! Not .txt format!")
		return
	}
	maski.ReadWholeFile(args[0])
}

func CheckTxt(s string) bool {
	if len(s) < 4 {
		return true
	}
	rev := myReverse(s)
	if rev[:4] != "txt." {
		return true
	}
	return false
}

func myReverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
