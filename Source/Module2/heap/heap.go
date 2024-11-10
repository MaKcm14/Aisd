package heap

import (
	"bufio"
	"errors"
	"fmt"
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

func (mheap *MinHeap) Search(key int) string {
	if index, isExist := mheap.lookup[key]; isExist {
		return fmt.Sprintf("1 %d %s", index, mheap.heap[index].value)
	}
	return "0"
}

func (mheap *MinHeap) Min() string {
	if len(mheap.heap) != 0 {
		return fmt.Sprintf("%d %d %s", mheap.heap[0].key, 0, mheap.heap[0].value)
	}
	return "error"
}

func (mheap *MinHeap) Max() string {
	if len(mheap.heap) == 0 {
		return "error"
	}

	maxIndex := 0
	for i := 1; i < len(mheap.heap); i++ {
		if mheap.heap[i].key > mheap.heap[maxIndex].key {
			maxIndex = i
		}
	}

	return fmt.Sprintf("%d %d %s", mheap.heap[maxIndex].key, maxIndex, mheap.heap[maxIndex].value)
}

func (mheap *MinHeap) Extract() string {
	if len(mheap.heap) == 0 {
		return "error"
	}
	root := mheap.heap[0]

	mheap.swap(0, len(mheap.heap)-1)
	mheap.heap = mheap.heap[:len(mheap.heap)-1]

	delete(mheap.lookup, root.key)
	if len(mheap.heap) > 0 {
		mheap.heapifyDown(0)
	}

	return fmt.Sprintf("%d %s", root.key, root.value)
}

func (mheap *MinHeap) Print() {
	if len(mheap.heap) == 0 {
		fmt.Println("_")
		return
	}

	fmt.Printf("[%v %v]\n", mheap.heap[0].key, mheap.heap[0].value)

	level, lastLevel, i := 2, 0, 1
	for interm := 2; i < len(mheap.heap); i++ {
		fmt.Printf("[%d %s %d] ", mheap.heap[i].key, mheap.heap[i].value, mheap.heap[(i-1)/2].key)
		if i == level {
			fmt.Print("\n")
			lastLevel = level
			level = level + 2*interm
			interm *= 2
		}
	}

	if (i - 1) != lastLevel {
		for ; i <= level; i++ {
			fmt.Print("_ ")
		}
		fmt.Println()
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

func Main() {
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
			fmt.Println(heap.Search(key))

		} else if ops[0] == "print" {
			heap.Print()

		} else if ops[0] == "min" {
			fmt.Println(heap.Min())

		} else if ops[0] == "max" {
			fmt.Println(heap.Max())

		} else if ops[0] == "extract" {
			fmt.Println(heap.Extract())

		} else {
			fmt.Println("error")
		}
	}
}
