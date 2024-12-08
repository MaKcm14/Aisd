package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

/*
Сложность алгоритма:
- Оценим сложность алгоритма по времени:
	- Сортировка: функция sort.Slice использует алгоритм быстрой сортировки => сложность её
		оценивается как O(nlogn).

	- Проход по массиву с попытками входа: O(n).

	- Итого: O(n + nlogn) = O(max{n, nlogn}) = O(nlogn), если не считать считывание ввода,
		которое так же O(n) и в целом не дает "внушительного вклада" в оценку,
		так как nlogn более быстрорастущая функция.

- Оценим сложность алгоритма по памяти:
	- По сути, происходит работа с массивом, переданным в функцию. Так как при оценке сложности по памяти
		оцениваем именно по количеству дополнительной памяти, необходимой для выполнения алгоритма,
		то алгоритм имеет сложность O(1) по памяти, ведь кроме стандартных переменных ему больше
		ничего не нужно.
*/

func BlockUser(loginTime []int, maxCountLogin, loginPeriod, blockTime, maxBlockTime, curTime int) {
	if len(loginTime) == 0 {
		fmt.Println("ok")
		return
	}

	var (
		timeUnblock int
		flagBlock   = false
		checkBlock  = false
	)

	sort.Slice(loginTime, func(i, j int) bool {
		return loginTime[i] < loginTime[j]
	})

	for i := 0; i < len(loginTime)-maxCountLogin+1; i++ {
		if timeUnblock <= loginTime[i+maxCountLogin-1] && flagBlock {
			flagBlock = false
		}

		if loginTime[i+maxCountLogin-1]-loginTime[i] <= loginPeriod {
			if checkBlock {
				blockTime = int(math.Min(float64(blockTime*2), float64(maxBlockTime)))
			}
			checkBlock = true
			flagBlock = true
			timeUnblock = loginTime[i+maxCountLogin-1] + blockTime
			i += maxCountLogin - 1
		}
	}

	if timeUnblock <= curTime {
		fmt.Println("ok")
	} else {
		fmt.Println(timeUnblock)
	}
}

func main() {
	var (
		maxCountLogin, loginPeriod, blockTime, maxBlockTime, curTime int
		userLoginTime                                                = make([]int, 0, 500)
	)

	fmt.Scan(&maxCountLogin, &loginPeriod, &blockTime, &maxBlockTime, &curTime)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var input = scanner.Text()
		time, err := strconv.Atoi(input)

		if err == nil && (time >= curTime-2*maxBlockTime || curTime <= 2*maxBlockTime) {
			userLoginTime = append(userLoginTime, time)
		}
	}

	BlockUser(userLoginTime, maxCountLogin, loginPeriod, blockTime, maxBlockTime, curTime)
}
