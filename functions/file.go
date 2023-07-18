package maski

import (
	"fmt"
	"io/ioutil"
)

func ReadWholeFile(fileName string) []float64 {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(string(contents)) == 0 {
		fmt.Println("Error! Empty file!")
		return nil
	}
	res := MainFunc(string(contents))
	return res
}
