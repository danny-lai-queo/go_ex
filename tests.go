package main

import (
	"container/list"
	"fmt"
	"log"
	"time"
)

func test_maze() {
	// This is how your code will be called.
	// Your answer should be the number of paths with the given cost.
	// You can edit this code to try different testing cases.
	cost := 8
	maze := NewMaze([][]int{
		{1, 2, 1},
		{6, 1, 1},
		{4, 3, 3},
	})
	learnerResult := countPaths(maze, 0, 0, cost)
	fmt.Println(learnerResult)
}

func test_cache() {
	// This is how your code will be called.
	// Your answer should contain the newest key value pairs.
	// You can edit this code to try different testing cases.
	keys := []Key{"potato", "broccoli", "carrot", "banana", "potato", "banana"}
	values := []Value{"0.45", "1.5", "0.75", "2", "0.35", "2.4"}
	n := NewCache()
	WriteValues(n, keys, values)
	learnerResult, learnerError := ReadValues(n, keys)
	if learnerError != nil {
		fmt.Println(learnerError.Error())
	} else {
		fmt.Println(learnerResult)
	}
}

func test_rolling_mean() {
	numbers := []int64{5, 7, 9, 6, 8, 10}
	input := make(chan int64)
	output := make(chan string)
	movingAvg := NewMovingAverage(input, output)
	go movingAvg.RollingMean()
	go func() {
		for _, n := range numbers {
			input <- n
		}
		close(input)
	}()

	learnerResult := ReadResults(output)
	fmt.Println(learnerResult)
}

func test_gr() {
	go func() {
		fmt.Println("hellow world.")
	}()
	time.Sleep(2000)
	func() {
		fmt.Println("aaaa aaaa aaaa aaaa aaaa aaaa aaaa aaaa aaaa aaaa ")
	}()
	time.Sleep(2000)
}

func test_chan() {
	numbers := []int{0, 1, 2, 3}
	results := make(chan int)
	for n := range results {
		go func(n int) {
			results <- n * n
		}(n)
	}

	for i := 0; i < len(numbers); i++ {
		log.Printf(" %d = %d\n", i, <-results)
	}
}

func test3() {
	dl := &DoublyLinkedList[string]{}
	if err := dl.Add(0, "Anton"); err == nil {
		dl.PrintFF()
	} else {
		fmt.Println(err.Error())
	}
}

func test4() {
	dl := &DoublyLinkedList[string]{}
	if err := dl.Add(99, "Anton"); err == nil {
		dl.PrintFF()
	} else {
		fmt.Println("Error:", err.Error())
	}
}

func test1() {
	// This is how your code will be called.
	// Your answer should respect the linked list order.
	// You can edit this code to try different testing cases.
	testCases := []struct {
		index int
		value string
	}{
		{index: 0, value: "C"},
		{index: 0, value: "A"},
		{index: 1, value: "B"},
		{index: 3, value: "D"},
	}
	dl := &DoublyLinkedList[string]{}
	// err := dl.AddElements(testCases)
	// forwardPrint := dl.PrintForward()
	// reversePrint := dl.PrintReverse()
	dl.AddElements(testCases)
	dl.PrintFF()
	dl.PrintRR()
}

func test2() {
	fmt.Println("testing")

	list1 := list.New()

	for i := 0; i < 9; i++ {
		list1.PushBack(i)
	}
	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Printf("v = %d\n", e.Value)
	}

	fmt.Println("......")

	stack1 := Stack[int]{}

	stack1.Push(7)
	stack1.Push(77)

	fmt.Printf("stack1 = %v\n", stack1)

	foo, _ := stack1.Pop()

	fmt.Printf("popped = %d\n", *foo)

	fmt.Printf("stack1 = %v\n", stack1)

	foo, _ = stack1.Pop()

	fmt.Printf("popped = %d\n", *foo)

	foo, err := stack1.Pop()

	fmt.Printf("foo = %d, err = %v\n", foo, err)

	fmt.Println("_________________")

	q := DoublyLinkedList[int]{}

	for i := 0; i < 5; i++ {
		if err_q := q.Add(i, i+10); err_q != nil {
			fmt.Printf("error %v\n", err_q)
		}
		q.Printd()
		fmt.Println("~~~~~~")
	}

	for i := 0; i < 5; i++ {
		if err_q := q.Add(i+1, i+20); err_q != nil {
			fmt.Printf("error %v\n", err_q)
		}
		q.Printd()
		fmt.Println("~~~~~~")
	}

	fmt.Printf("q.size = %d\n", q.size)
	fmt.Printf("q.head = %v\n", q.head)
	fmt.Printf("q.tail = %v\n", q.tail)

	fmt.Printf("%s\n", q.PrintForward())
	fmt.Printf("%s\n", q.PrintReverse())

}
