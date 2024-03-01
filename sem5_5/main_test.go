package main

import "testing"

func TestAdd(t *testing.T) {
	exp := 5
	x := 3
	y := 2

	res := add(x, y)

	if res != exp {
		t.Fail()
	}

	exp = 5
	x = 2
	y = 2

	res = add(x, y)

	if res != exp {
		t.Logf("4 not 5")
		//go test -v -count=1 .
		t.Fail()

	}
}

// go test -v -count=1 -run TestAddV2
func TestAddV2(t *testing.T) {
	for _, v := range []struct {
		x   int
		y   int
		exp int
	}{
		{
			x:   1,
			y:   2,
			exp: 3,
		},
		{
			x:   2,
			y:   2,
			exp: 4,
		},
		{
			x:   5,
			y:   2,
			exp: 7,
		},
		{
			x:   5,
			y:   2,
			exp: 8,
		},
	} {
		res := add(v.x, v.y)
		if res != v.exp {
			t.Logf("%v не равно %v\n", res, v.exp)
			t.Fail()
		}
	}
}
