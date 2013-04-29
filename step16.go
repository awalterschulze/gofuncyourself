//Ideas stolen from rob pike's example before methods could be assigned like functions
//https://code.google.com/p/go/source/browse/src/pkg/sort/example_keys_test.go?spec=svn2bf8d07c14c178039b3220beb5d76fc07b56358a&name=2bf8d07c14c1&r=2bf8d07c14c178039b3220beb5d76fc07b56358a

//cast functions: words, forward, reverse to type By
//call their Sort methods
//Sort creates a struct with a function member
package main

import (
	"fmt"
	"sort"
)

type Question struct {
	Words string
	Votes int
}

type By func(q1, q2 *Question) bool

func (by By) Sort(questions []Question) {
	ps := &questionSorter{
		questions: questions,
		by:        by,
	}
	sort.Sort(ps)
}

type questionSorter struct {
	questions []Question
	by        func(p1, p2 *Question) bool
}

func (s *questionSorter) Len() int {
	return len(s.questions)
}

func (s *questionSorter) Swap(i, j int) {
	s.questions[i], s.questions[j] = s.questions[j], s.questions[i]
}

func (s *questionSorter) Less(i, j int) bool {
	return s.by(&s.questions[i], &s.questions[j])
}

func main() {
	questions := []Question{
		Question{"We have seen this before, why are you doing this?", 2},
		Question{"Do something useful", 1},
		Question{"Go Func yourself", 10},
	}

	words := func(q1, q2 *Question) bool {
		return q1.Words < q2.Words
	}
	forward := func(q1, q2 *Question) bool {
		return q1.Votes < q2.Votes
	}
	reverse := func(q1, q2 *Question) bool {
		return q1.Votes > q2.Votes
	}

	fmt.Printf("before %v\n", questions)
	By(forward).Sort(questions)
	fmt.Printf("sorted %v\n", questions)
	By(reverse).Sort(questions)
	fmt.Printf("reverse %v\n", questions)
	By(words).Sort(questions)
	fmt.Printf("alphabetical %v\n", questions)
}
