// ビジネスロジックはここに書く

package stringer

import "strconv"

func Reverse(input string) (result string) {
	// 入力文字列を逆転させる
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func Inspect(input string, digits bool) (count int, kind string) {
	// 入力値が文字列か数字かを検査する
	if !digits {
		return len(input), "char"
	}
	return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}
