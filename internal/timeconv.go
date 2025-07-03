package internal

import (
	"fmt"
	"strconv"
	"time"
)

// ConvertTimestamp 将unix timestamp转换为指定格式的日期时间字符串
func ConvertTimestamp(timestampStr string) (string, error) {
	// 解析timestamp字符串为整数
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return "", fmt.Errorf("无效的timestamp格式: %v", err)
	}

	// 根据timestamp的位数判断是秒还是毫秒
	var t time.Time
	if timestamp > 9999999999 { // 大于10位数，假设是毫秒
		t = time.Unix(timestamp/1000, (timestamp%1000)*1000000)
	} else { // 10位数或更少，假设是秒
		t = time.Unix(timestamp, 0)
	}

	// 格式化为指定格式 (YYYY-MM-DD HH:MM:SS)
	return t.Format("2006-01-02 15:04:05"), nil
}

// ConvertDateToTimestamp 将日期时间字符串转换为unix timestamp
func ConvertDateToTimestamp(dateStr string, outputMillis bool) (string, error) {
	// 支持的日期格式
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
	}

	var t time.Time
	var err error

	// 尝试不同的日期格式
	for _, format := range formats {
		t, err = time.Parse(format, dateStr)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", fmt.Errorf("无效的日期格式: %v, 支持的格式: YYYY-MM-DD HH:MM:SS, YYYY-MM-DD HH:MM, YYYY-MM-DD", err)
	}

	// 返回秒级或毫秒级时间戳
	if outputMillis {
		return strconv.FormatInt(t.Unix()*1000, 10), nil
	}
	return strconv.FormatInt(t.Unix(), 10), nil
}
