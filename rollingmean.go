// Write your answer here, and then test your code.
// Your job is to implement the RollingMean() method.

package main

import (
	"fmt"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
// const showExpectedResult = false
// const showHints = false

const WindowSize = 3

type NumericType interface {
	int64 | float64
}

type MovingAverage[T NumericType] struct {
	Position     int
	WindowValues [WindowSize]T
	WindowFilled bool
	Input        <-chan T
	Output       chan<- string
}

func NewMovingAverage[T NumericType](in <-chan T,
	out chan<- string) *MovingAverage[T] {
	return &MovingAverage[T]{
		Input:  in,
		Output: out,
	}
}

// CalculateMean prints the window and its calculated mean.
func (ma *MovingAverage[T]) CalculateMean() string {
	values := make([]string, WindowSize)
	var sum float64
	for i, v := range ma.WindowValues {
		values[i] = fmt.Sprint(v)
		sum += float64(v)
	}
	mean := sum / WindowSize
	return fmt.Sprintf("[%s] = %.2f", strings.Join(values, ","), mean)
}

func (ma *MovingAverage[T]) RollingMean() {
	for {
		theInput, ok := <-ma.Input

		if !ok {
			close(ma.Output)
			return
		}

		cpIndex := ma.Position % WindowSize

		ma.WindowValues[cpIndex] = theInput

		if ma.WindowFilled {
			ma.Output <- ma.CalculateMean()
		}

		ma.Position += 1

		if ma.Position == (WindowSize - 1) {
			ma.WindowFilled = true
		}
	}
}

func (ma *MovingAverage[T]) RollingMean_old() {
	//fmt.Println("RollingMean not implemented")
	for {
		theInput, ok := <-ma.Input

		//if _, ok := <-ma.Input; !ok {
		if !ok {
			if ma.WindowFilled {
				ma.Output <- ma.CalculateMean()
			}
			close(ma.Output)
			return
		} else {
			//fmt.Printf("theInput = %v\n", theInput)
			if !ma.WindowFilled {
				ma.WindowValues[ma.Position] = theInput
				ma.Position += 1
				ma.WindowFilled = (ma.Position == WindowSize)
			} else {
				ma.Output <- ma.CalculateMean()
				cpIndex := ma.Position % WindowSize
				ma.WindowValues[cpIndex] = theInput
				ma.Position += 1
			}

			//fmt.Printf("WindowValues = %v\n", ma.WindowValues)
		}
	}
}

func ReadResults(output <-chan string) []string {
	var results []string
	for r := range output {
		results = append(results, r)
	}
	return results
}
