package other

import "fmt"

// Boy 男孩
type Boy struct {
}

// Say 说话
func (b *Boy) Say(s string) {
	fmt.Println(s)
}

// Inv 增加
func Inv() (v int) {
	defer func() { v++ }()
	return 42
}
