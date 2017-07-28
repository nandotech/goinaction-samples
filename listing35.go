// Sample program to show how a bytes.Buffer can also be used
// with the io.Copy function
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main. Entry point
func main() {
	var b bytes.Buffer

	//Write a string to the buffer
	b.Write([]byte("Hello"))

	//Use Fprintf to concatenate a string to the buffer
	fmt.Fprintf(&b, "World!")

	//Write the content of the buffer to stdout
	// buffer, err :=
	io.Copy(os.Stdout, &b)
	// if err != nil {
	// 	fmt.Println("Error %v", err)
	// 	return
	// }
	// fmt.Println(buffer, err)
}
