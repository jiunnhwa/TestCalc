package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestCalculator(t *testing.T) {
	g := Goblin(t)
	g.Describe("Calculator", func() {
		g.It("should add two numbers ", func() {
			g.Assert(Add(1, 2)).Equal(3)
		})

		g.It("should subtract two numbers", func() {
			g.Assert(Subtract(5, 3)).Equal(2) 
		})

		g.It("should multiply two numbers", func() {
			g.Assert(Multiply(5, 6)).Equal(30)
		})

		g.It("should divide two numbers", func() {
			g.Assert(Divide(10, 2)).Equal(float64(5))
		})
		
		g.It("should divide negative numbers", func() {  //New Test.
			g.Assert(Divide(4, -2)).Equal(float64(-2))
		})		
	})
}
