package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	fn *object.CompileFunction
	ip int
}

func NewFrame(fn *object.CompileFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
