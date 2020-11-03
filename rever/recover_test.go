package rever

import (
	"log"
	"testing"
)

func TestHanppen(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("发生异常", r)
		}
	}()
	Hanppen()
}
