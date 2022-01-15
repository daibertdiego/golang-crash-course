package strutil

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
