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

type BloomFilter struct {
	buckets     []byte
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

	return BloomFilter{
		apprSize:    uint64(n),
		possibility: p,
		m:           uint64(m),
		k:           uint64(k),
		buckets:     make([]byte, m),
		primes:      make([]uint64, 0, k),
	}, nil
}

func (filter *BloomFilter) calcPrimes() {
	j, i := uint64(3), uint64(1)

	if len(filter.primes) == 0 {
		filter.primes = append(filter.primes, 2)
	} else {
		j = filter.primes[len(filter.primes)-1] + 1
		i = uint64(len(filter.primes))
	}

	for ; i < filter.k; j++ {
		if res := filter.isPrime(j); res {
			filter.primes = append(filter.primes, j)
			i++
		}
	}
}

func (filter *BloomFilter) isPrime(num uint64) bool {
	for i := uint64(2); i <= uint64(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func (filter *BloomFilter) Add(key uint64) error {
	if filter.buckets == nil {
		return errors.New("try to add the element to uninitialized bloom filter")
	}

	filter.calcPrimes()

	for i := uint64(0); i != filter.k; i++ {
		bucketNum := (((i + 1) % MersenNum31) * (key % MersenNum31)) % MersenNum31
		bucketNum += filter.primes[i] % MersenNum31
		bucketNum %= MersenNum31
		bucketNum %= filter.m

		filter.buckets[bucketNum] = 1
	}

	return nil
}

func (filter *BloomFilter) Search(key uint64) (bool, error) {
	if filter.buckets == nil {
		return false, errors.New("try to search the element in the uninitialized bloom filter")
	}
	filter.calcPrimes()

	for i := uint64(0); i != filter.k; i++ {
		bucketNum := (((i + 1) % MersenNum31) * (key % MersenNum31)) % MersenNum31
		bucketNum += filter.primes[i] % MersenNum31
		bucketNum %= MersenNum31
		bucketNum %= filter.m

		if res := filter.buckets[bucketNum]; res == 0 {
			return false, nil
		}
	}
	return true, nil
}

func (filter *BloomFilter) Print(w io.Writer) error {
	if filter.buckets == nil {
		return errors.New("try to print the unitialized bloom filter")
	}

	for _, elem := range filter.buckets {
		fmt.Fprint(w, elem)
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
