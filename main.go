package main

import (
	"log"
)

func main() {
	workingDir, err := utils.CurrentDir()
	if err != nil {
		log.Println("Unable to fetch current dir,%s\n", err.Error())
	}

	execDir, err := utils.TerraformPath()
	if err != nil {
		log.Println("Unable to execute dir:%s\n", err.Error())
	}
	cli.InitAndExecute(workingDir, execDir)
}
