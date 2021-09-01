package bijection

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	l501 := Make[int, string]()

	l501.Set(7567, "Rex")
	l501.Set(5555, "Fives")
	l501.Set(5597, "Jesse")
	l501.Set(408, "Echo")

	printEach := func(id int, name string) bool {
		fmt.Println(id, name)
		return true
	}

	l501.Range(printEach)

	fmt.Println()

	l501.DeleteX(5597)
	l501.DeleteY("Echo")

	inv := l501.Inverse()

	inv.Range(SwapRangeFunc[int, string](printEach))
}
