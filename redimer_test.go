package redimer

import (
	"reflect"
	"testing"
)

func TestArray1(t *testing.T) {
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := Redim(&arr, 3, 3)
	want := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v:", want, got)
	}
}

func TestArray2(t *testing.T) {
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := Redim(&arr, 2, 4)
	want := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v:", want, got)
	}
}

func TestSlice1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := Redim(&arr, 3, 3)
	want := [][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v:", want, got)
	}
}

func TestSlice2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := Redim(&arr, 2, 4)
	want := [][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v:", want, got)
	}
}
