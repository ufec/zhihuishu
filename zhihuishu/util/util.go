package util

import (
	"fmt"
	"strconv"
	"time"
)

// GetStringTimestamp 获取13位时间戳
func GetStringTimestamp() string {
	return strconv.FormatInt(time.Now().Unix() * 1000, 10)
}

// SecToTim 秒数转为时:分:秒
func SecToTim(sec int) string {
	hour := sec / 3600
	sec %= 3600
	minute := sec / 60
	second := sec % 60
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}