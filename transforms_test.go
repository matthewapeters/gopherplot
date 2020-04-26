package gopherplot_test

import (
	"fmt"
	"testing"

	gp "github.com/matthewapeters/gopherplot"
)

func TestVector(t *testing.T) {
	v, err := gp.NewVector([]int{5, 1}, 1, 3, 6, 1, 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
	v, err = gp.NewVector([]int{1, 5}, 1, 3, 6, 1, 2)
	if err != nil {
		t.Error(err)
	}
	v.Plus(&gp.Number{6})
	fmt.Println(v)
	v.Times(&gp.Number{2})
	fmt.Println(v)

	z := v.T()
	fmt.Println(z)

	// Changing the internal value of a vector will change the internal value of its transpose
	z.Data[1].Val = -100.00
	if z.Data[1].Value() != v.Data[1].Value() {
		t.Error("Expected values to be 100.00000")
	}

	x, err := z.DotProduct(v)
	if err != nil {
		t.Error(err)
	}
	if x.Value() != 11224.0 {
		t.Errorf("Expected 11224.0, got %f", x.Value().(float64))
	}

	two, err := gp.NewVector([]int{1, 1}, 2)
	if err != nil {
		t.Error(err)
	}
	twoDotTwo, err := two.DotProduct(two.T())
	if err != nil {
		t.Error(err)
	}
	if twoDotTwo.Value() != 4.0 {
		t.Error("Expected [2.0] . [2.0] to be 4.0")
	}
}

func TestMatrix(t *testing.T) {
	m, err := gp.NewMatrix([]int{3, 3}, 1, 0, 0, 0, 100, 0, 0, 1, 1)
	if err != nil {
		t.Error(err)
	}
	x, err := m.Get(1, 1)
	if err != nil {
		t.Error(err)
	}
	if x.Value().(float64) != 100.0 {
		t.Errorf("Expected 100.0000, got %f", x)
	}
	m.AddNumber(&gp.Number{200})
	x, err = m.Get(1, 1)
	if err != nil {
		t.Error(err)
	}
	if x.Value().(float64) != 300.0 {
		t.Error(err)
	}
	m.AddMatrix(m)
	x, err = m.Get(1, 1)
	if err != nil {
		t.Error(err)
	}
	if x.Value().(float64) != 600.0 {
		t.Error(err)
	}
}
