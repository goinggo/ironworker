// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This package implements the WorkManager singleton. This manager
	controls the starting, shutdown and processing of work.
*/
package program

import (
	"github.com/goinggo/ironio/helper"
	"time"
)

func DoWork(shutChan chan bool, waitChan chan bool, logKey string) {

	helper.WriteStdout("Program", "program.DoWork", "Program Started")

	defer func() {

		waitChan <- true
	}()

	// Perform work for 10 seconds
	for count := 0; count < 40; count++ {

		select {

		case <-shutChan:

			helper.WriteStdout("Program", "program.DoWork", "Info : Completed : KILL REQUESTED")
			shutChan <- true
			return

		default:

			helper.WriteStdoutf("Program", "program.DoWork", "Info : Performing Work : %d", count)
			time.Sleep(time.Millisecond * 250)
		}
	}

	helper.WriteStdout("Program", "program.DoWork", "Completed")
}
