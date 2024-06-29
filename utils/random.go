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

// f:需要处理的浮点数，n：要保留小数的位数
// Pow10（）返回10的n次方，最后一位四舍五入，对ｎ＋１位加０．５后四舍五入
func RoundFloat(f float64, n int) float64 {
	n10 := math.Pow10(n)
	if f < 0 {
		return math.Trunc((math.Abs(f)+0.5/n10)*n10*-1) / n10
	} else {
		return math.Trunc((math.Abs(f)+0.5/n10)*n10) / n10
	}
}
