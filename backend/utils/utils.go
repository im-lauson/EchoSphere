package utils

import "strings"

// TODO: 添加通用的工具函数

// 判断字符串是否为空
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
