package generator

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func FirstNamesGen(quantity uint64) (firstNames []string) {
	names := []string{"鈴木", "佐々木", "小林", "中村", "吉田", "石塚", "竹田", "原田", "阿部", "伊藤", "西村", "白川", "茂木", "小野寺", "岸田", "金子", "青木", "広野", "上田", "目黒"}

	return randomStrings(names, quantity)
}

func LastNamesGen(quantity uint64) (lastNames []string) {
	names := []string{"聡", "祐一", "徹", "和幸", "敬一郎", "博文", "俊一", "裕子", "ひとみ", "洋子", "智子", "千春", "勇気", "夏樹", "祥", "美由紀", "晃", "博信"}

	return randomStrings(names, quantity)
}

func MiddleNamesGen(quantity uint64) (middleNames []string) {
	names := []string{"Blue", "Joan of Arc", "Lil", "August", "March", "April"}
	return randomStrings(names, quantity)
}

func FullNamesGen(quantity uint64) (middleNames []string) {
	fullNames := make([]string, quantity)

	f := FirstNamesGen(quantity)
	l := LastNamesGen(quantity)

	for i := range fullNames {
		fullNames[i] = f[i] + l[i]
	}

	return fullNames
}

func NumberGen(quantity uint64, baselineNumber int64, marginOfError int64, allowMinusNum bool) (nums []int64) {
	var i uint64
	var n int64

	for i = 0; i < quantity; i++ {
		for {
			n = baselineNumber + (rand.Int63n(2*marginOfError) - marginOfError)
			if allowMinusNum || n >= 0 {
				break
			}
		}
		nums = append(nums, n)
	}
	return nums
}

func randomStrings(strs []string, quantity uint64) (retStrs []string) {
	retStrs = make([]string, quantity)
	var i uint64

	for i = 0; i < quantity; i++ {
		str := strs[rand.Intn(len(strs))]
		retStrs = append(retStrs, str)
	}
	return retStrs
}
