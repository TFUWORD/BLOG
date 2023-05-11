package utils

// 判断字符串key是否在列表list里面
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
