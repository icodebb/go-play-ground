/**
 * Test struct and functions.
 */

package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Operator is role.
type Operator struct {
	Name     string
	Age      int
	workTime int
}

// MyLog is my wrapper of logrus.
func MyLog(level int, msg ...interface{}) {
	line := fmt.Sprintf("%v", msg...)
	log.Info(line)
}

// MyLogf is my wrapper of logrus with format.
func MyLogf(level int, format string, msg ...interface{}) {
	line := fmt.Sprintf(format, msg...)
	log.Infof(line)
}

// SimpleTest tests log and Fibonacci.
func SimpleTest() {
	// Test function and struct
	op := new("Ice", 46, 24)
	t, err := op.mySum()
	if err == nil {
		log.Infof("Years before work: %v", t)
	}

	// Test numbers
	for n := 0; n < 10; n++ {
		log.Infof("%v", Fibonacci(n))
	}
}

func new(name string, age int, working int) *Operator {
	op := &Operator{
		Name:     "Ice Liu",
		Age:      46,
		workTime: 24,
	}

	return op
}

func (c *Operator) myOwnAdd(x int, y int) (int, error) {
	z := c.Age + c.workTime
	return z, nil
}

func (c *Operator) mySum() (int, error) {
	return c.Age - c.workTime, nil
}
