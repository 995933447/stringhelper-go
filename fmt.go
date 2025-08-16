package stringhelper

import "strings"

func Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	var j bool
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		// 判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	// ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

func Camel(s string) string {
	data := make([]byte, 0, len(s))
	num := len(s) - 1
	var j bool
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		if j {
			d = d - 32
			j = false
		}
		data = append(data, d)
	}
	return string(data)
}

func LowerFirstASCII(s string) string {
	if s == "" {
		return s
	}
	b := s[0]
	if 'A' <= b && b <= 'Z' {
		// 小写化：大写字母和小写字母的差值是 32
		return string(b+('a'-'A')) + s[1:]
	}
	return s
}
