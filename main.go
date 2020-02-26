package main

import (
	"fmt"
	"sort"
)

type Task struct {
	id int
	time int
	start int
	priority int
}

type State int;

const (
	n State = 0
	r State = 1
	p State = 2
)

type Process struct {
	id int
	val []State
}


type Algo struct {};

func (a *Algo) FCFS(tasks []Task) {
	// sort by start INT moment

	var prs []Process;

	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].start < tasks[j].start;
	})

	for i := 0; i < len(tasks); i++ {

	}

	fmt.Println(tasks, prs);
}

func main() {
	fmt.Println(n, r, p);
	sol := &Algo{};

	sol.FCFS([]Task{
		{
			id:1,
			time:3,
			start:2,
		},
		{
			id:2,
			time:4,
			start:0,
		},
		{
			id:3,
			time:5,
			start:6,
		},
	})
}