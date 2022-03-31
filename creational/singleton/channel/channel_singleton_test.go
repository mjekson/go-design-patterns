package channel_singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()
	singleton3 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
		go singleton3.AddOne()
	}
	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*3 {
		val = singleton.GetCount()
		fmt.Printf("After loop, current count is %d\n", val)
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()
}
