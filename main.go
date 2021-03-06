/**
 * MyPlayGround project entry.
 */

package main

import (
	"sync"

	"github.com/icodebb/go-play-ground/ch"
	"github.com/icodebb/go-play-ground/dt"
	"github.com/icodebb/go-play-ground/menu"
	"github.com/icodebb/go-play-ground/num"
	"github.com/icodebb/go-play-ground/utils"
	log "github.com/sirupsen/logrus"
)

// Define a type of function for map.
type fn func()

// Wait for your goroutines to finish.
var wg sync.WaitGroup

func main() {
	// Only 世界 works on Linux.
	// fmt.Printf("⌘ and %v 世界\n", "\U00002714")

	m := map[int]fn{
		0: utils.MyVersion,
		1: SimpleTest,
		2: num.NumTest,
		3: dt.TestTime,
		4: ch.TestChennel,
	}

	log.Infoln("Start")

	// Waiting for user's input.
	for {
		r := menu.PrintMenu()

		if r != 99 {
			m[r]()
		} else {
			return
		}
	}

	// wg.Add(1)
	// ExampleWithCancel(&wg)
	// wg.Wait()

	// Test JSON
	// TestJson0()

	log.Infoln("Done")
}
