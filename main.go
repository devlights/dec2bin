package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments\tusage:: dec2bin dec-value(e.g: 255)")
		return 1
	}

	v := args[0]
	b, err := Convert(v)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return 2
	}

	fmt.Printf("%s\t-->\t%s\n", v, b)

	return 0
}

// Convert -- 指定された10進数文字列を2進数に変換します.
func Convert(v string) (string, error) {
	if len(v) == 0 {
		return "", nil
	}

	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("0b%b", i), nil
}
