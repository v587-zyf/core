package utils

import (
	"math"
	"math/rand"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RoundFloat 使用 math.Round 将浮点数 f 四舍五入到小数点后 n 位。
func RoundFloat(f float64, n int) float64 {
	if f < 0 {
		return f
	}
	shift := math.Pow(10, float64(n))
	// 将浮点数乘以10的n次方，四舍五入到最近的整数，然后再除以10的n次方。
	return math.Round(f*shift) / shift
}
