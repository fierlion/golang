package bitset

import (
	"bytes"
	"fmt"
)

//an BitSet is a set of small non-negative integers.
//its zero value represents the empty set.
type BitSet struct {
	words []uint64
	size int
}

//Has reports whether the set containst the non-negative value x
func (s *BitSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//Add adds the non-negative value x to the set
func (s *BitSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	s.size += 1
}

//Remove clears the non-negative value x from the set
func (s *BitSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^(1<< bit)
	s.size -= 1
}

//Clear sets the bitset to 0
func (s *BitSet) Clear() {
	s.words = make([]uint64, 0)
	s.size = 0
}

//UnionWith sets s to the union of s and t.
func (s *BitSet) UnionWith(t *BitSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//Len returns the size of the bitset
func (s *BitSet) Len() int {
	return s.size
}

//Copy returns a deep copy of the BitSet
func (s *BitSet) Copy() *BitSet {
  cpy := BitSet{s.words, s.size}
	return &cpy
}

//String returns the set as a string of the form "{1 2 3}"
func (s *BitSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
