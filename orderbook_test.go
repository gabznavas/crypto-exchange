package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)

	buyOrderA := NewOrder(true, 0.2)
	buyOrderB := NewOrder(true, 0.5)
	buyOrderC := NewOrder(true, 0.7)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	fmt.Printf("%+v\n", l)

	if len(l.Orders) != 2 {
		t.Fatalf("%+v", l)
	}
	if l.Orders[0] != buyOrderA {
		t.Fatalf("%+v", l)
	}
	if l.Orders[1] != buyOrderC {
		t.Fatalf("%+v", l)
	}
}

func TestOrderbook(t *testing.T) {

}
