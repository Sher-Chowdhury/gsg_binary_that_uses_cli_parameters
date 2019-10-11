package main

// need to first run: go get github.com/fatih/color
import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	myMessage := ""
	arguments := os.Args

	fmt.Println(reflect.TypeOf(arguments)) // this is a slice
	// https://golang.org/pkg/builtin/#len
	if len(arguments) == 1 {
		myMessage = "No arguements provided"
		fmt.Println(myMessage)
		// here's how to exit out of a binary early with an exit code.
		os.Exit(1)
	} else {
		myMessage = arguments[1]
	}

	// https://golang.org/pkg/io
	// io.WriteString() works in a similar way to fmt.Print(). the main difference is that fmt.Print() can take multiple parameters, e.g. fmt.Print(var1, "some stuff", var2, "even more stuff", "etc")
	// wheerase io.WriteString() can only take exactly 2 params, where to write to, and what to write.
	io.WriteString(os.Stdout, myMessage)
	io.WriteString(os.Stdout, "\n")

}
