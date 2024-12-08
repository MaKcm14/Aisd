package knapsack

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	Weight int
	Cost   int
}

type Store struct {
	SumWeight   int
	SumApprCost int
	SumCost     int
	ItemNums    []int
}

func Knapsack(approx float64, maxPrice int, items map[int]*Item, maxWeight int) (Store, error) {
	if len(items) == 0 {
		return Store{}, errors.New("empty items' slice was given")
	}

	coef := approx * float64(maxPrice) / float64(len(items))
	maxCost := 0
	stores := make(map[int]Store)
	stores[0] = Store{
		ItemNums:    []int{},
		SumWeight:   0,
		SumApprCost: 0,
	}

	for num, item := range items {
		lastStores := make(map[int]Store)

		for key, val := range stores {
			lastStores[key] = val
		}

		for _, store := range lastStores {
			sumWeight := store.SumWeight + item.Weight
			sumCost := store.SumApprCost + int(math.Floor(float64(item.Cost)/coef))

			if oldStore, flagExist := stores[sumCost]; (store.SumWeight+item.Weight <= maxWeight) &&
				(sumWeight < oldStore.SumWeight || !flagExist) {
				if maxCost < sumCost {
					maxCost = sumCost
				}

				stores[sumCost] = Store{
					ItemNums:    append(append(make([]int, 0, len(store.ItemNums)+1), store.ItemNums...), num),
					SumWeight:   sumWeight,
					SumApprCost: sumCost,
				}
			}
		}
	}

	store := stores[maxCost]
	for _, num := range stores[maxCost].ItemNums {
		store.SumCost += items[num].Cost
	}

	sort.Slice(store.ItemNums, func(i, j int) bool {
		return store.ItemNums[i] < store.ItemNums[j]
	})

	return store, nil
}

func Main() {
	var (
		approx     float64
		maxWeight  int
		count      int
		curItemNum = 1
		maxCost    = 0
		items      = make(map[int]*Item, 500)
	)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if len(input) != 0 && count == 0 {
			approx, _ = strconv.ParseFloat(input, 32)
			count++
		} else if len(input) != 0 && count == 1 {
			maxWeight, _ = strconv.Atoi(input)
			count++

		} else if len(input) != 0 {
			item := strings.Split(input, " ")

			weight, _ := strconv.Atoi(item[0])

			price, _ := strconv.Atoi(item[1])

			if maxWeight >= weight {
				items[curItemNum] = &Item{
					Weight: weight,
					Cost:   price,
				}

				if price > maxCost {
					maxCost = price
				}
			}
			curItemNum++
		}
	}
	res, err := Knapsack(approx, maxCost, items, maxWeight)

	if err != nil {
		fmt.Println(0, 0)
		return
	}

	fmt.Println(res.SumWeight, res.SumCost)
	for _, val := range res.ItemNums {
		fmt.Println(val)
	}
}
