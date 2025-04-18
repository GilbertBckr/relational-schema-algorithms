package set

import (
	"fmt"
	"sort"
)

type Set struct {
	items map[string]bool
}

func New() *Set {

	Set := &Set{
		items: make(map[string]bool),
	}

	return Set
}
func NewFromElements(inputs []string) *Set {

	set := &Set{
		items: make(map[string]bool),
	}
	for _, e := range inputs {
		set.Add(e)
	}

	return set
}

func (s *Set) Contains(item string) bool {
	res, ok := s.items[item]

	if !ok {
		return false
	}

	return res
}

func (s *Set) Add(item string) {
	s.items[item] = true
}

func (s *Set) Remove(item string) {
	delete(s.items, item)
}

func (s *Set) IsSubSet(s2 *Set) bool {
	for key, value := range s.items {
		if !value {
			continue
		}

		if !s2.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set) GetElementsOrdered() []string {
	// TODO: surely there is a more efficient way
	keys := make([]string, 0, len(s.items))

	for k, v := range s.items {
		if !v {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys

}

func (s *Set) DeepCopy() *Set {
	set := New()

	for k, v := range s.items {
		if !v {
			continue
		}
		set.Add(k)
	}

	return set
}

func (s *Set) Equal(s2 *Set) bool {
	// There is a more efficient way but this is cool because it closely mimick the mathematical definition for Set equality
	return s.IsSubSet(s2) && s2.IsSubSet(s)
}

// Modifies the set by adding the elements of s2 to the current set
func (s *Set) AddUnion(s2 *Set) {
	for k, v := range s2.items {
		if !v {
			continue
		}
		s.Add(k)
	}

}

func (s *Set) string() string {
	return fmt.Sprintf("%v", s.GetElementsOrdered())
}
