package main

import (
	"os"
	"fmt"
	"github.com/Sameetpatro/odlang/lexer"
)

func main(){
	fmt.Println("OS ARGS")
	for i := 0; i < len(os.Args); i++{
		fmt.Println(os.Args[i])
	}
	if len(os.Args) < 2 {
		println("Usage: jagannath <file.od>")
		return
	}
	src, err := os.ReadFile(os.Args[1])
	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	l := lexer.New(string(src))
	l.Debug()
}