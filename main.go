package main

import (
	"github.com/goinggo/ironworker/controller"
	"github.com/goinggo/ironworker/helper"
)

func main() {

	helper.WriteStdout("Main", "main", "Started")

	controller.Run()

	helper.WriteStdout("Main", "main", "Completed")
}
