package set

import (
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	set := New()

	set.add("hello")

	if !set.contains("hello") {
		t.Errorf("Set does not contain 'hello'")
	}
}

func TestRemoveItem(t *testing.T) {
	set := New()

	set.add("hello")

	set.remove("hello")

	if set.contains("hello") {
		t.Errorf("Set does not contain 'hello'")
	}
}

func TestEmptySubset(t *testing.T) {
	set1 := New()
	set2 := New()

	if !set1.isSubset(set2) {
		t.Fatalf("Empty subset is not subset of empty set")
	}
}

func TestNonEmptySubsetOfEmpty(t *testing.T) {
	empty := New()
	set2 := New()
	set2.add("A")

	if !empty.isSubset(set2) {
		t.Fatalf("empty is not subset of {A}")
	}
	if set2.isSubset(empty) {
		t.Fatalf("empty is not subset of {A}")
	}
}
func TestEqualNonEmpty(t *testing.T) {
	set1 := New()
	set1.add("A")
	set1.add("B")
	set2 := New()
	set2.add("A")
	set2.add("B")

	if !set1.isSubset(set2) {
		t.Fatalf("{A, B} is subset of {A, B}")
	}
	if !set2.isSubset(set1) {
		t.Fatalf("{A, B} is subset of {A, B}")
	}
}

func TestTrueSubset(t *testing.T) {
	set1 := New()
	set1.add("A")
	set2 := New()
	set2.add("A")
	set2.add("B")

	if !set1.isSubset(set2) {
		t.Fatalf("{A} is subset of {A, B}")
	}
	if set2.isSubset(set1) {
		t.Fatalf("{A, B} is not subset of {A}")
	}
}

func TestOrderedItems(t *testing.T) {
	set := New()

	if len(set.getElementsOrdered()) != 0 {
		t.Fatalf("Slice should be empty")
	}

	set.add("A")

	if !reflect.DeepEqual(set.getElementsOrdered(), []string{"A"}) {
		t.Fatal("Sorted elements are not equal to [A]")
	}

	set.add("C")
	set.add("B")

	if !reflect.DeepEqual(set.getElementsOrdered(), []string{"A", "B", "C"}) {
		t.Fatal("Sorted elements are not equal to [A,B,C]")
	}

}
