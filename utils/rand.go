package utils

import (
	"math/rand"
)

/**
*随机权重
*randData map[int]int{索引:权重，索引：权重}
*return 索引
 */
func RandWeightByMap(randData map[int]int) int {
	sum := 0
	for _, v := range randData {
		sum += v
	}
	if sum <= 0 {
		return -1
	}
	randNum := rand.Intn(sum)
	count := 0
	for k, v := range randData {
		count += v
		if randNum < count {
			return k
		}
	}
	return -1
}
