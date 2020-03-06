/**
 * MyPlayGround project entry.
 */

package main

import (
	"sync"

	utils "github.com/icodebb/go-play-ground/utils"
	log "github.com/sirupsen/logrus"
)

// Wait for your goroutines to finish.
var wg sync.WaitGroup

func main() {
	log.Infoln("Start")

	utils.MyVersion()

	log.Infof("Random int:%v, string:%s", RandomInt(1, 10), RandomString(8))

	// ch := make(chan int)

	//SimpleTest()

	// MyLog(1, "Logging")
	// MyLogf(1, "I am %v.", "logging")

	wg.Add(1)
	ExampleWithCancel(&wg)
	wg.Wait()

	// Test time
	// TestTime()

	// Test JSON
	// TestJson0()

	// Test Chennels
	// TestChennel()

	log.Infoln("Done")

	// fatal error: all goroutines are asleep - deadlock!
	// fails if no subroutines
	// log.Infof("ch = %v", <-ch)
}
