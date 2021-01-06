package kits2

import (
	"fmt"
	"testing"
)

func Test_counterkit(t *testing.T) {
	c := NewCounterKit("b")
	c.Inc()

	fmt.Println(c.Show())

}
