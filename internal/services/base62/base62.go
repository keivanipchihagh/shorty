package base62

const (
	base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Encode(num int64) string {
	var result string
	for num > 0 {
		mod := num % 62
		num /= 62
		result = string(base62[mod]) + result
	}
	return result
}

func Decode(str string) int64 {
	var result int64
	for _, char := range str {
		result *= 62
		result += int64(char - '0')
	}
	return result
}
