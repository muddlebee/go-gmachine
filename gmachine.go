// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

type Word int

// DefaultMemSize is the default size of the memory of a G-machine	.
const DefaultMemSize = 1024

// Gmachine is a virtual CPU.
type Gmachine struct {
	P      Word
	Memory []Word
}

func New() *Gmachine {
	return &Gmachine{
		P:      0,
		Memory: make([]Word, DefaultMemSize),
	}
}
