package main

func try(word string, ToFind string) bool {
	for i := 0; i <= len(ToFind)-len(word); i++ {
		if ToFind[i:i+len(word)] == ToUpper(word) || ToFind[i:i+len(word)] == ToLower(word) {
			return true
		}
	}
	return false
}

func ToUpper(s string) string {
	temp := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) >= "a" && string(s[i]) <= "z" {
			temp += string(s[i] - 32)
		} else {
			temp += string(s[i])
		}
	}
	return temp
}

func ToLower(s string) string {
	temp := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) >= "A" && string(s[i]) <= "Z" {
			temp += string(s[i] + 32)
		} else {
			temp += string(s[i])
		}
	}
	return temp
}
