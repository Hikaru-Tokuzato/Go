package main

import (
	"fmt"
	"time"
)

const (
	core = 4
	N    = 10000
)

var (
	list       [N]int = [N]int{}
	sosuuList         = []int{}
	num        int
	i          int = 0
	sosuuJudge     = true
	baisuu     int
)

func makeList() {
	for i := 0; i < N; i++ {
		list[i] = i + 1
		//fmt.Println(sosuu[i])
	}
	sosuuList = list[:]
}

func remove(ints []int, search int) []int {
	result := []int{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

func hurui() {
	num = sosuuList[i]
	if num == 1 {
		sosuuList = remove(sosuuList, 1)
	} else {
		if num != sosuuList[len(sosuuList)-1] {
			for j := 1; j <= N/num; j++ {
				if j != 1 {
					sosuuList = remove(sosuuList, num*j)
				}
			}
			i++
		} else {
			sosuuJudge = false
		}
	}

}

func kakeru(a int, b int) int {

	baisuu = a * b

	fmt.Printf("%d＊%d＝%d\n", a, b, baisuu)

	return baisuu
}

func yosoku(baisuu int) {

	for i := 0; i < len(sosuuList); i++ {
		if baisuu%sosuuList[i] == 0 {
			fmt.Println(sosuuList[i])
			fmt.Println(baisuu / sosuuList[i])
			break
		}
	}

}

func main() {

	start := time.Now()

	makeList()

	for sosuuJudge {
		hurui()
		fmt.Println(num)
	}

	fmt.Println(sosuuList)

	end := time.Now()

	/*

		rand.Seed(time.Now().UnixNano())

		var a int = rand.Intn(len(sosuuList))

		var b int = rand.Intn(len(sosuuList))

		kakeru(sosuuList[a], sosuuList[b])

		yosoku(baisuu)

	*/

	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
