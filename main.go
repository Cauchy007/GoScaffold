package main

import "fmt"

import "github.com/Cauchy007/GoScaffold/cmd"

func main() {
	fmt.Println("Hello go-scaffold")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
