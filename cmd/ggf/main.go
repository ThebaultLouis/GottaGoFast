package main

import (
	"fmt"
	"github.com/ThebaultLouis/GottaGoFast/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
