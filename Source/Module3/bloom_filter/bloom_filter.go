package bloom

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	MersenNum31 uint64 = uint64(math.Pow(2, 31)) - 1
)

type BitArray struct {
	array []byte
}

func NewBitArray(size int64) BitArray {
	return BitArray{
		array: make([]byte, (size>>3)+1),
	}
}

func (arr *BitArray) Set(index uint64) {
	byteNum := index / 8
	elemPos := index % 8

	setNum := byte(math.Pow(2, float64(7-elemPos)))
	checkSet := arr.array[byteNum] & setNum

	if checkSet == 0 {
		arr.array[byteNum] += setNum
	}
}

func (arr *BitArray) Get(index uint64) int {
	byteNum := index / 8
	elemPos := index % 8

	setNum := byte(math.Pow(2, float64(7-elemPos)))
	checkSet := arr.array[byteNum] & setNum

	if checkSet == 0 {
		return 0
	}

	return 1
}

func (arr *BitArray) IsInit() bool {
	return !(arr.array == nil)
}

type BloomFilter struct {
	buckets     BitArray
	apprSize    uint64
	m           uint64
	k           uint64
	possibility float64
	primes      []uint64
}

func NewBloomFilter(n int64, p float64) (BloomFilter, error) {
	m := -int64(math.Round(float64(n) * math.Log2(p) / math.Ln2))
	k := -int64(math.Round(math.Log2(p)))

	if n <= 0 || p <= 0 || m <= 0 || k <= 0 {
		return BloomFilter{}, errors.New("try to create filter with the wrong parameters")
	}

	primes := calcPrimes()

	return BloomFilter{
		apprSize:    uint64(n),
		possibility: p,
		m:           uint64(m),
		k:           uint64(k),
		buckets:     NewBitArray(m),
		primes:      primes,
	}, nil
}

func calcPrimes() []uint64 {
	startNum := uint64(2)
	primes := make([]uint64, 0)
	nums := make([]uint64, 0, 1000)

	for i := uint64(0); i != 1000; i++ {
		nums = append(nums, startNum+i)
	}

	curNum := startNum
	for curNum*curNum < 1000 {
		for i := 2*curNum - startNum; i < uint64(len(nums)); i += curNum {
			if nums[i]%curNum == 0 {
				nums[i] = 0
			}
		}

		for i := uint64(curNum-startNum) + 1; i < uint64(len(nums)); i++ {
			if nums[i] != 0 {
				curNum = nums[i]
				break
			}
		}
	}

	for _, val := range nums {
		if val != 0 {
			primes = append(primes, val)
		}
	}

	return primes
}

func (filter *BloomFilter) Add(key uint64) error {
	if !filter.buckets.IsInit() {
		return errors.New("try to add the element to uninitialized bloom filter")
	}

	for i := uint64(0); i != filter.k; i++ {
		bucketNum := (((i + 1) % MersenNum31) * (key % MersenNum31)) % MersenNum31
		bucketNum += filter.primes[i] % MersenNum31
		bucketNum %= MersenNum31
		bucketNum %= filter.m

		filter.buckets.Set(bucketNum)
	}

	return nil
}

func (filter *BloomFilter) Search(key uint64) (bool, error) {
	if !filter.buckets.IsInit() {
		return false, errors.New("try to search the element in the uninitialized bloom filter")
	}

	for i := uint64(0); i != filter.k; i++ {
		bucketNum := (((i + 1) % MersenNum31) * (key % MersenNum31)) % MersenNum31
		bucketNum += filter.primes[i] % MersenNum31
		bucketNum %= MersenNum31
		bucketNum %= filter.m

		if res := filter.buckets.Get(bucketNum); res == 0 {
			return false, nil
		}
	}
	return true, nil
}

func (filter *BloomFilter) Print(w io.Writer) error {
	if !filter.buckets.IsInit() {
		return errors.New("try to print the unitialized bloom filter")
	}

	for i := uint64(0); i != filter.m; i++ {
		fmt.Fprint(w, filter.buckets.Get(i))
	}
	fmt.Fprintln(w)

	return nil
}

func main() {
	var filter BloomFilter
	var flagInit = false
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		commds := strings.Split(input, " ")

		if commds[0] == "set" && len(commds) == 3 {
			if flagInit {
				fmt.Println("error")
				continue
			}
			n, _ := strconv.ParseInt(commds[1], 10, 64)
			p, _ := strconv.ParseFloat(commds[2], 64)

			f, err := NewBloomFilter(n, p)

			filter = f
			if err != nil {
				fmt.Println("error")
			} else {
				flagInit = true
				fmt.Println(filter.m, filter.k)
			}

		} else if commds[0] == "add" && len(commds) == 2 {
			key, _ := strconv.ParseUint(commds[1], 10, 64)

			err := filter.Add(uint64(key))

			if err != nil {
				fmt.Println("error")
			}

		} else if commds[0] == "search" && len(commds) == 2 {
			key, _ := strconv.ParseUint(commds[1], 10, 64)

			flagFind, err := filter.Search(uint64(key))

			if err != nil {
				fmt.Println("error")
			} else if flagFind {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}

		} else if commds[0] == "print" && len(commds) == 1 {
			err := filter.Print(os.Stdout)

			if err != nil {
				fmt.Println("error")
			}
		}
	}
}
