package main

import (
	"fmt"
	"github.com/lincaiyong/arg"
	"os"
	"strings"
)

func main() {
	os.Args = strings.Fields("./test --input=test.txt --output=test2.txt --force xxxx sdfsdf")
	arg.Parse()
	fmt.Println(arg.Args())
	fmt.Println(arg.BoolArgs())
	fmt.Println(arg.KeyValueArgs())
}
