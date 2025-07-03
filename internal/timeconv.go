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
