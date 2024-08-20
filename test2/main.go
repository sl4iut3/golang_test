package main

import (
	"fmt"
	"sort"
)

type ListeTriable []int

func (a ListeTriable) Len() int           { return len(a) }
func (a ListeTriable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ListeTriable) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	done := make(chan bool)
	var l = []int{5, 6, 3, 2}
	//var l = ListeTriée([]int{5, 6, 3, 2})
	sort.Sort(ListeTriable(l))

	fmt.Println(l)

	for _, v := range l {
		go func(int) {
			fmt.Printf("%d ", v)
			done <- true
		}(0)
	}
	for range l {
		<-done
	}
}
