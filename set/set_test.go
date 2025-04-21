package set

import (
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	set := New()

	set.Add("hello")

	if !set.Contains("hello") {
		t.Errorf("Set does not contain 'hello'")
	}
}

func TestRemoveItem(t *testing.T) {
	set := New()

	set.Add("hello")

	set.Remove("hello")

	if set.Contains("hello") {
		t.Errorf("Set does not contain 'hello'")
	}
}

func TestEmptySubset(t *testing.T) {
	set1 := New()
	set2 := New()

	if !set1.IsSubSet(set2) {
		t.Fatalf("Empty subset is not subset of empty set")
	}
}

func TestNonEmptySubsetOfEmpty(t *testing.T) {
	empty := New()
	set2 := New()
	set2.Add("A")

	if !empty.IsSubSet(set2) {
		t.Fatalf("empty is not subset of {A}")
	}
	if set2.IsSubSet(empty) {
		t.Fatalf("empty is not subset of {A}")
	}
}
func TestEqualNonEmpty(t *testing.T) {
	set1 := New()
	set1.Add("A")
	set1.Add("B")
	set2 := New()
	set2.Add("A")
	set2.Add("B")

	if !set1.IsSubSet(set2) {
		t.Fatalf("{A, B} is subset of {A, B}")
	}
	if !set2.IsSubSet(set1) {
		t.Fatalf("{A, B} is subset of {A, B}")
	}
}

func TestTrueSubset(t *testing.T) {
	set1 := New()
	set1.Add("A")
	set2 := New()
	set2.Add("A")
	set2.Add("B")

	if !set1.IsSubSet(set2) {
		t.Fatalf("{A} is subset of {A, B}")
	}
	if set2.IsSubSet(set1) {
		t.Fatalf("{A, B} is not subset of {A}")
	}
}

func TestOrderedItems(t *testing.T) {
	set := New()

	if len(set.GetElementsOrdered()) != 0 {
		t.Fatalf("Slice should be empty")
	}

	set.Add("A")

	if !reflect.DeepEqual(set.GetElementsOrdered(), []string{"A"}) {
		t.Fatal("Sorted elements are not equal to [A]")
	}

	set.Add("C")
	set.Add("B")

	if !reflect.DeepEqual(set.GetElementsOrdered(), []string{"A", "B", "C"}) {
		t.Fatal("Sorted elements are not equal to [A,B,C]")
	}

}

func TestIsEmpty(t *testing.T) {
	set := New()

	if !set.IsEmpty() {
		t.Fatalf("Set was not marked as empty")
	}

	set.Add("A")

	if set.IsEmpty() {
		t.Fatalf("Set was not marked as not empty")
	}

	set.Remove("A")

	if !set.IsEmpty() {
		t.Fatalf("Set was not marked as empty")
	}

}
