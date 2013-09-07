// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This package controls the running and termination of the program.
	Used to test if IronWorker will send OS signals on a call to kill.
*/
package controller

import (
	"fmt"
	"github.com/goinggo/ironworker/helper"
	"github.com/goinggo/ironworker/program"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	NAMESPACE  = "controller"
	GO_ROUTINE = "main"
)

// Run is the entry point for the controller
func Run() {

	helper.WriteStdout("Main", "controller.Run", "Started")

	// Create a channel to talk with the OS
	sigChan := make(chan os.Signal, 1)

	// Create a channel to let the program tell us it is done
	waitChan := make(chan bool)

	// Create a channel to shut down the program early
	shutChan := make(chan bool)

	// Launch the work routine
	go program.DoWork(shutChan, waitChan, "Test")

	// Ask the OS to notify us about events
	signal.Notify(sigChan)

	for {

		select {

		case whatSig := <-sigChan:

			// Convert the signal to an integer so we can display the hex number
			sigAsInt, _ := strconv.Atoi(fmt.Sprintf("%d", whatSig))

			helper.WriteStdoutf("Main", "controller.Run", "******> OS Notification: %v : %#x", whatSig, sigAsInt)

			// Did we get any of these termination events
			if whatSig == syscall.SIGINT ||
				whatSig == syscall.SIGKILL ||
				whatSig == syscall.SIGQUIT ||
				whatSig == syscall.SIGSTOP ||
				whatSig == syscall.SIGTERM {

				helper.WriteStdout("Main", "controller.Run", "******> Program Being Killed")

				// Close the signal channel to tell the program to shutdown
				shutChan <- true
				<-shutChan

				helper.WriteStdout("Main", "controller.Run", "******> Shutting Down")
				return
			}

			continue

		case <-waitChan:

			helper.WriteStdout("Main", "controller.Run", "******> Shutting Down")
			return
		}
	}

	helper.WriteStdout("Main", "controller.Run", "Completed")
}
