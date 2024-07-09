package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// MD5 md5加密
func MD5(src string) string {
	w := md5.New()
	w.Write([]byte(src))
	return hex.EncodeToString(w.Sum(nil))
}

// GUID 产生新的GUID
func GUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	strMD5 := base64.URLEncoding.EncodeToString(b)
	return MD5(strMD5)
}

// Token 产生新的用户登录验证码
func Token() string {
	return fmt.Sprint(time.Now().Unix(), ":", GUID())
}

func Int32ArrayToString(src []int32, flag string) (out string) {
	// 没有的就直接返回
	if len(src) == 0 {
		return ""
	}
	out = ""
	for k, v := range src {
		if k == len(src)-1 {
			out = fmt.Sprint(out, v)
		} else {
			out = fmt.Sprint(out, v, flag)
		}
	}

	return
}

func StringToInt32Array(src, flag string) (out []int32) {
	// 没有的就直接返回
	if src == "" {
		return nil
	}

	strs := strings.Split(src, flag)
	for _, v := range strs {
		data, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}

		out = append(out, int32(data))
	}

	return
}

func StringArrayToInt32Array(src []string) (out []int32) {
	for _, v := range src {
		data, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		out = append(out, int32(data))
	}

	return
}

func StrToInt32(src string) int32 {
	data, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}

	return int32(data)
}

func StrToFloat(src string) float64 {
	data, err := strconv.ParseFloat(src, 32)
	if err != nil {
		return 0
	}

	return data
}

func StrToInt64(src string) int64 {
	data, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return 0
	}

	return data
}

func StrToUInt64(src string) uint64 {
	data, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return 0
	}

	return uint64(data)
}

func StrToInt(src string) int {
	data, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}

	return data
}
