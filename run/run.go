package run

import (
	"io"

	"github.com/Nv7-Github/bpp/ir"
)

type Runnable struct {
	Stdout io.Writer
	Index  int

	ir        *ir.IR
	registers []interface{}
	vars      map[int]interface{}

	params []int
	args   []string
}

func NewRunnable(ir *ir.IR) *Runnable {
	return &Runnable{
		ir:        ir,
		registers: make([]interface{}, len(ir.Instructions)),
		vars:      make(map[int]interface{}),
	}
}

func (r *Runnable) Run(args []string) error {
	r.args = args
	for i := range r.ir.Instructions {
		r.Index = i
		if err := r.runInstruction(i); err != nil {
			return err
		}
	}
	return nil
}
