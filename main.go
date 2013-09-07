// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This program tests using the Iron.io IronWorker product.

	Ardan Studios
	12973 SW 112 ST, Suite 153
	Miami, FL 33186
	bill@ardanstudios.com

	GoingGo.net Post:
	http://www.goinggo.net/2013/07/running-go-programs-in-ironworker.html
*/
package main

import (
	"github.com/goinggo/ironworker/controller"
	"github.com/goinggo/ironworker/helper"
)

// main is the entry point for the program
func main() {

	helper.WriteStdout("Main", "main", "Started")

	controller.Run()

	helper.WriteStdout("Main", "main", "Completed")
}
