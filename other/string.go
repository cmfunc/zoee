package other

// 直接比较时，从左至右逐个字符根据ASCII码值大小比较
func StringDiff() bool {
	return "Alice" > "Bob"
}