package main

import (
	"github.com/goinggo/ironio/controller"
	"github.com/goinggo/ironio/helper"
)

func main() {

	helper.WriteStdout("Main", "main", "Started")

	controller.Run()

	helper.WriteStdout("Main", "main", "Completed")
}
