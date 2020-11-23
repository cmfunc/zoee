package alg

import (
	"log"
	"testing"
)

func TestBubbleSort(t *testing.T){
	log.Println(BubbleSort([]int{90,3,800,2,34,4,13,6,10,2,45,11}))
}