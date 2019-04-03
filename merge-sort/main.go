package main

import (
	"fmt"
	"sync"
	"time"
)

// MergeSort ...
func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := s[:n]
	r := s[n:]

	lRet := MergeSort(l)
	rRet := MergeSort(r)

	return merge(lRet, rRet)
}

var sem = make(chan struct{}, 5)

func MergeSortMulti(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	wg := sync.WaitGroup{}
	wg.Add(2)
	var l []int
	var r []int
	go func() {
		l = MergeSortMulti(s[:n])
		wg.Done()
	}()
	go func() {
		r = MergeSortMulti(s[n:])
		wg.Done()
	}()

	wg.Wait()
	return merge(l, r)
}

func merge(a []int, b []int) []int {
	ret := []int{}
	for len(a) > 0 || len(b) > 0 {
		if len(a) == 0 {
			ret = append(ret, b...)
			break
		}
		if len(b) == 0 {
			ret = append(ret, a...)
			break
		}
		if a[0] <= b[0] {
			ret = append(ret, a[0])
			a = a[1:]
		} else {
			ret = append(ret, b[0])
			b = b[1:]
		}
	}
	return ret
}

func MergeSortMultiEffi(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	var l = []int{}
	var r = []int{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	n := len(s) / 2
	select {
	case sem <- struct{}{}:
		go func() {
			l = MergeSortMultiEffi(s[:n])
			<-sem
			wg.Done()
		}()
	default:
		l = MergeSortMultiEffi(s[:n])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			r = MergeSortMultiEffi(s[n:])
			<-sem
			wg.Done()
		}()
	default:
		r = MergeSortMultiEffi(s[n:])
		wg.Done()
	}
	wg.Wait()
	return merge(l, r)
}

func main() {
	fmt.Println("------------")

	t1 := time.Now().UnixNano()
	fmt.Println(MergeSort([]int{2, 4, 3, 6, 2, 7}))
	t2 := time.Now().UnixNano()
	fmt.Println("time consumed by mergeSort:=", t2-t1)

	t1 = time.Now().UnixNano()
	fmt.Println(MergeSortMulti([]int{2, 4, 3, 6, 2, 7}))
	t2 = time.Now().UnixNano()
	fmt.Println("time consumed by mergeSortMulti:=", t2-t1)

}
