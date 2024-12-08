package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MinHeap struct {
	heap   []Element
	lookup map[int]int
}

type Element struct {
	key   int
	value string
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap:   make([]Element, 0, 200),
		lookup: make(map[int]int, 200),
	}
}

func (mheap *MinHeap) Add(key int, value string) error {
	if _, isExist := mheap.lookup[key]; isExist {
		return errors.New("try to add the existing key")
	}
	mheap.heap = append(mheap.heap, Element{key, value})
	mheap.lookup[key] = len(mheap.heap) - 1
	mheap.heapifyUp(len(mheap.heap) - 1)
	return nil
}

func (mheap *MinHeap) Set(key int, value string) error {
	if index, isExist := mheap.lookup[key]; isExist {
		mheap.heap[index].value = value
		return nil
	}
	return errors.New("try to set the value to unexisting key")
}

func (mheap *MinHeap) Delete(key int) error {
	if index, isExist := mheap.lookup[key]; isExist {
		mheap.swap(index, len(mheap.heap)-1)
		mheap.heap = mheap.heap[:len(mheap.heap)-1]

		delete(mheap.lookup, key)

		if index < len(mheap.heap) {
			mheap.heapifyUp(index)
			mheap.heapifyDown(index)
		}

		return nil
	}
	return errors.New("try to delete the element with the unexisting key")
}

func (mheap *MinHeap) Search(key int) (int, string, bool) {
	if index, isExist := mheap.lookup[key]; isExist {
		return index, mheap.heap[index].value, true
	}
	return -1, "", false
}

func (mheap *MinHeap) Min() (int, string, bool) {
	if len(mheap.heap) != 0 {
		return mheap.heap[0].key, mheap.heap[0].value, true
	}
	return -1, "", false
}

func (mheap *MinHeap) Max() (int, string, int, bool) {
	if len(mheap.heap) == 0 {
		return -1, "", -1, false
	}

	maxIndex := 0

	for i := len(mheap.heap) / 2; i != len(mheap.lookup); i++ {
		if mheap.heap[i].key > mheap.heap[maxIndex].key {
			maxIndex = i
		}
	}

	return mheap.heap[maxIndex].key, mheap.heap[maxIndex].value, maxIndex, true
}

func (mheap *MinHeap) Extract() (int, string, bool) {
	if len(mheap.heap) == 0 {
		return -1, "", false
	}
	root := mheap.heap[0]

	mheap.swap(0, len(mheap.heap)-1)
	mheap.heap = mheap.heap[:len(mheap.heap)-1]

	delete(mheap.lookup, root.key)
	if len(mheap.heap) > 0 {
		mheap.heapifyDown(0)
	}

	return root.key, root.value, true
}

func (mheap *MinHeap) Print(w io.Writer) {
	if len(mheap.heap) == 0 {
		fmt.Fprintln(w, "_")
		return
	}

	fmt.Fprintf(w, "[%v %v]\n", mheap.heap[0].key, mheap.heap[0].value)

	level, lastLevel, i := 2, 0, 1
	for interm := 2; i < len(mheap.heap); i++ {
		fmt.Fprintf(w, "[%d %s %d] ", mheap.heap[i].key, mheap.heap[i].value, mheap.heap[(i-1)/2].key)
		if i == level {
			fmt.Fprintln(w)
			lastLevel = level
			level = level + 2*interm
			interm *= 2
		}
	}

	if (i - 1) != lastLevel {
		for ; i <= level; i++ {
			fmt.Fprint(w, "_ ")
		}
		fmt.Fprintln(w)
	}
}

func (mheap *MinHeap) heapifyUp(i int) {
	for i > 0 && mheap.heap[i].key < mheap.heap[(i-1)/2].key {
		mheap.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (mheap *MinHeap) heapifyDown(i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i

	if left < len(mheap.heap) && mheap.heap[left].key < mheap.heap[smallest].key {
		smallest = left
	}

	if right < len(mheap.heap) && mheap.heap[right].key < mheap.heap[smallest].key {
		smallest = right
	}

	if smallest != i {
		mheap.swap(i, smallest)
		mheap.heapifyDown(smallest)
	}
}

func (mheap *MinHeap) swap(i, j int) {
	mheap.heap[i], mheap.heap[j] = mheap.heap[j], mheap.heap[i]
	mheap.lookup[mheap.heap[i].key] = i
	mheap.lookup[mheap.heap[j].key] = j
}

func parseInput(input string) ([]string, error) {
	var commds = strings.Split(input, " ")

	if len(commds) != 0 && (((commds[0] == "add" || commds[0] == "set") && len(commds) == 3) ||
		((commds[0] == "delete" || commds[0] == "search") && len(commds) == 2) ||
		((commds[0] == "print" || commds[0] == "min" || commds[0] == "max" || commds[0] == "extract") &&
			len(commds) == 1)) {
		return commds, nil
	} else {
		return nil, fmt.Errorf("error")
	}
}

func main() {
	var heap = NewMinHeap()
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var input = scanner.Text()

		if input == "" {
			continue
		}
		ops, err := parseInput(input)

		if err != nil {
			ops = []string{"error"}
		}

		if ops[0] == "add" {
			key, _ := strconv.Atoi(ops[1])
			err := heap.Add(key, ops[2])

			if err != nil {
				fmt.Println("error")
			}

		} else if ops[0] == "set" {
			key, _ := strconv.Atoi(ops[1])
			err := heap.Set(key, ops[2])

			if err != nil {
				fmt.Println("error")
			}

		} else if ops[0] == "delete" {
			key, _ := strconv.Atoi(ops[1])
			err := heap.Delete(key)

			if err != nil {
				fmt.Println("error")
			}

		} else if ops[0] == "search" {
			key, _ := strconv.Atoi(ops[1])

			index, val, flagFind := heap.Search(key)
			if flagFind {
				fmt.Printf("1 %d %s\n", index, val)
			} else {
				fmt.Println("0")
			}

		} else if ops[0] == "print" {
			heap.Print(os.Stdout)

		} else if ops[0] == "min" {
			key, val, flagFind := heap.Min()

			if flagFind {
				fmt.Printf("%d %d %s\n", key, 0, val)
			} else {
				fmt.Println("error")
			}

		} else if ops[0] == "max" {
			key, val, index, flagFind := heap.Max()

			if flagFind {
				fmt.Printf("%d %d %s\n", key, index, val)
			} else {
				fmt.Println("error")
			}

		} else if ops[0] == "extract" {
			key, val, flagFind := heap.Extract()

			if flagFind {
				fmt.Printf("%d %s\n", key, val)
			} else {
				fmt.Println("error")
			}

		} else {
			fmt.Println("error")
		}
	}
}
