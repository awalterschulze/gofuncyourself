//you're a reverse
package main

import (
	"fmt"
	"sort"
)

type Question struct {
	Words string
	Votes int
}

type LessFunc func(i, j int) bool

type Questions struct {
	List []Question
	LessFunc
}

func (this *Questions) Len() int {
	return len(this.List)
}

func (this *Questions) Less(i, j int) bool {
	return this.LessFunc(i, j)
}

func (this *Questions) Swap(i, j int) {
	this.List[i], this.List[j] = this.List[j], this.List[i]
}

func (this *Questions) Forward(i, j int) bool {
	return this.List[i].Votes < this.List[j].Votes
}

func (this *Questions) Reverse(i, j int) bool {
	return this.List[i].Votes > this.List[j].Votes
}

func main() {
	questions := &Questions{
		[]Question{
			Question{"We have seen this before, why are you doing this?", 2},
			Question{"Do something useful", 1},
			Question{"Go Func yourself", 10},
		},
		nil,
	}
	fmt.Printf("before %v\n", questions)
	questions.LessFunc = questions.Forward
	sort.Sort(questions)
	fmt.Printf("sorted %v\n", questions)
	questions.LessFunc = questions.Reverse
	sort.Sort(questions)
	fmt.Printf("reverse %v\n", questions)
}
