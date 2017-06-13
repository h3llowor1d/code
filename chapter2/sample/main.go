package main

import (
	"log"
	"os"
	"github.com/code/chapter2/sample/search"
	_ "github.com/code/chapter2/sample/matchers"
	"flag"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	var searchStr string
	flag.StringVar(&searchStr, "search","","搜索关键字")
	//解析输入的参数
	flag.Parse()
	// Perform the search for the specified term.
	search.Run(searchStr)
}
