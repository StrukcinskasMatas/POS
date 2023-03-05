package main

import (
	"pos/cmd"
	"pos/utils"
)

func main() {
	utils.InitializeLogger()
	cmd.Execute()
}
