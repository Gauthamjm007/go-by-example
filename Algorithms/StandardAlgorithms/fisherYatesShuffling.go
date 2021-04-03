package FisherYatesShuffling

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func FisherYatesShuffling(slice interface{}) {
	rand.Seed(time.Now().UnixNano())
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
	fmt.Println(slice)
}
