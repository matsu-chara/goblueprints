package trace

import (
	"fmt"
	"io"
)

// Tracer is Tracer
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// New crates Tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates SilentTracer
func Off() Tracer {
	return &nilTracer{}
}
