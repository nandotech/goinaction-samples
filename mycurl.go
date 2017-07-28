package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//init function that runs before main
func init() {
	if (len(os.Args)) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

//Entry point
func main() {
	//Get a response from web server
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	//Copies from body to stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
