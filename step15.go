//you're a reverse even by less
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

func (this LessFunc) Less(i, j int) bool {
	return this(i, j)
}

type Questions struct {
	List []Question
	LessFunc
}

func (this *Questions) Len() int {
	return len(this.List)
}

func (this *Questions) Swap(i, j int) {
	this.List[i], this.List[j] = this.List[j], this.List[i]
}

func (this *Questions) Reverse(i, j int) bool {
	return this.List[i].Votes > this.List[j].Votes
}

func (this *Questions) Forward(i, j int) bool {
	return this.List[i].Votes < this.List[j].Votes
}

func (this *Questions) Words(i, j int) bool {
	return this.List[i].Words < this.List[j].Words
}

func (this *Questions) By(L LessFunc) sort.Interface {
	this.LessFunc = L
	return this
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
	sort.Sort(questions.By(questions.Forward))
	fmt.Printf("sorted %v\n", questions)
	sort.Sort(questions.By(questions.Reverse))
	fmt.Printf("reverse %v\n", questions)
	sort.Sort(questions.By(questions.Words))
	fmt.Printf("alphabetical %v\n", questions)
}
