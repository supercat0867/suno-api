package kernel

import (
	"net/http"
	"strings"
	"time"
)

// parseCookies 解析 cookie 字符串为 http.Cookie 结构体切片
func parseCookies(cookieStr string) []*http.Cookie {
	cookiePairs := strings.Split(cookieStr, "; ")
	cookies := make([]*http.Cookie, len(cookiePairs))
	for i, cookiePair := range cookiePairs {
		parts := strings.SplitN(cookiePair, "=", 2)
		if len(parts) != 2 {
			continue
		}
		cookies[i] = &http.Cookie{
			Name:  parts[0],
			Value: parts[1],
		}
	}
	return cookies
}

// 将毫秒级时间戳转换为time.Time格式
func convertTimestampMillisToTime(timestamp int64) time.Time {
	// 将毫秒转换为秒和纳秒
	seconds := timestamp / 1000
	nanoseconds := (timestamp % 1000) * 1_000_000
	t := time.Unix(seconds, nanoseconds)
	return t
}
