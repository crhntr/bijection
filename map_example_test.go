package bijection

import "testing"

func TestMap(t *testing.T) {
	l501 := New(map[int]string{
		7567: "Rex",
		5555: "Fives",
		5597: "Jesse",
	})

	l501.Add(408, "Echo")
	if l501.Len() != 4 {
		t.Fail()
	}

	l501.DeleteA(5597)
	l501.DeleteB("Echo")
	if l501.Len() != 2 {
		t.Fail()
	}

	l501.Range(func(int, string) {})
	l501.Range(SwapRangeFunc(func(string, int) {}))

	inv := l501.Inverse()
	inv.Range(func(string, int) {})
	inv.Range(SwapRangeFunc(func(int, string) {}))
}
