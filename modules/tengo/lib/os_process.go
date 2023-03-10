package lib

import (
	"os"
	"syscall"

	"github.com/sin3degrees/tengo/v2"
	"github.com/sin3degrees/tengo/v2/stdlib"
)

func makeOSProcessState(state *os.ProcessState) *tengo.ImmutableMap {
	return &tengo.ImmutableMap{
		Value: map[string]tengo.Object{
			"exited": &tengo.UserFunction{
				Name:  "exited",
				Value: stdlib.FuncARB(state.Exited),
			},
			"pid": &tengo.UserFunction{
				Name:  "pid",
				Value: stdlib.FuncARI(state.Pid),
			},
			"string": &tengo.UserFunction{
				Name:  "string",
				Value: stdlib.FuncARS(state.String),
			},
			"success": &tengo.UserFunction{
				Name:  "success",
				Value: stdlib.FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *tengo.ImmutableMap {
	return &tengo.ImmutableMap{
		Value: map[string]tengo.Object{
			"kill": &tengo.UserFunction{
				Name:  "kill",
				Value: stdlib.FuncARE(proc.Kill),
			},
			"release": &tengo.UserFunction{
				Name:  "release",
				Value: stdlib.FuncARE(proc.Release),
			},
			"signal": &tengo.UserFunction{
				Name: "signal",
				Value: func(args ...tengo.Object) (tengo.Object, error) {
					if len(args) != 1 {
						return nil, tengo.ErrWrongNumArguments
					}
					i1, ok := tengo.ToInt64(args[0])
					if !ok {
						return nil, tengo.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &tengo.UserFunction{
				Name: "wait",
				Value: func(args ...tengo.Object) (tengo.Object, error) {
					if len(args) != 0 {
						return nil, tengo.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
