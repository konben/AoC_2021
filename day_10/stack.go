package main

type stack []byte

func (s *stack) push(val byte) {
	(*s) = append((*s), val)
}

func (s *stack) pop() byte {
	if len(*s) == 0 {
		panic("Cannot pop an empty stack!")
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}
