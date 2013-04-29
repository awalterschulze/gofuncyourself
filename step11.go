//let's look at sort again
package main

import (
	"fmt"
	"sort"
)

type Question struct {
	Words string
	Votes int
}

type Questions struct {
	List []Question
}

func (this *Questions) Len() int {
	return len(this.List)
}

func (this *Questions) Less(i, j int) bool {
	return this.List[i].Votes < this.List[j].Votes
}

func (this *Questions) Swap(i, j int) {
	this.List[i], this.List[j] = this.List[j], this.List[i]
}

func main() {
	questions := &Questions{
		[]Question{
			Question{"We have seen this before, why are you doing this?", 2},
			Question{"Do something useful", 1},
			Question{"Go Func yourself", 10},
		},
	}
	fmt.Printf("before %v\n", questions)
	sort.Sort(questions)
	fmt.Printf("after %v\n", questions)
}
