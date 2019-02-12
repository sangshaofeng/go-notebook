package utils

import (

)

// 检查用户名是否合法，只接受数字和字母，且第一位非数字
func CheckUsernme(username string) bool {
	if username[0] >= '0' && username[0] <= '9' {
		return false
	}
	for i := 0; i < len(username); i++ {
		if !(username[i] == '_' ||
			(username[i] >= '0' && username[i] <= '9') ||
			(username[i] >= 'a' && username[i] <= 'z') ||
			(username[i] >= 'A' && username[i] <= 'Z')) {
			return false
		}
	}
	return true
}