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

func (p *Process) appendToState (s State) []State {
	p.val = append(p.val, s);
	return p.val;
}

type filled struct {
	end int
}

func (f *filled) canBeFilled(n int) bool {
	a := n > f.end;
	return a;
}

type Algo struct {};

func (a *Algo) FCFS(tasks []Task) {
	// sort by start INT moment
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].start < tasks[j].start;
	})

	// iterate the processes
	prs := make([]Process, len(tasks));
	var lastEnd int = 0;
	for i := 0; i < len(tasks); i++ {
		prs[i].id = tasks[i].id;
		var pEnd int = tasks[i].start + tasks[i].time;
		var pStart int = tasks[i].start;

		//iterate and fill processes
		for j := 0; j < 20; j++ {
			if i == 0 {
				if j >= pStart && j <= pEnd {
					prs[i].appendToState(p)
				} else {
					prs[i].appendToState(n)
				}
			} else {
				if j >= pStart && j <= lastEnd {
					prs[i].appendToState(r);
					pEnd++;
				} else if j > lastEnd && j <= pEnd {
					prs[i].appendToState(p)
				} else {
					prs[i].appendToState(n)
				}
			}
		}
		lastEnd = pEnd;
	}

	fmt.Println(prs);
}

func main() {
	fmt.Println(n, r, p);
	sol := &Algo{};

	sol.FCFS([]Task{
		{
			id:1,
			time:3,
			start:1,
		},
		{
			id:2,
			time:4,
			start:3,
		},
		{
			id:3,
			time:5,
			start:6,
		},
	})
}