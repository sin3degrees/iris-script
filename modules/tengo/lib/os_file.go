package lib

import (
	"os"

	"github.com/sin3degrees/tengo/v2"
	"github.com/sin3degrees/tengo/v2/stdlib"
)

func makeOSFile(file *os.File) *tengo.ImmutableMap {
	return &tengo.ImmutableMap{
		Value: map[string]tengo.Object{
			// chdir() => true/error
			"chdir": &tengo.UserFunction{
				Name:  "chdir",
				Value: stdlib.FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &tengo.UserFunction{
				Name:  "chown",
				Value: stdlib.FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &tengo.UserFunction{
				Name:  "close",
				Value: stdlib.FuncARE(file.Close),
			}, //
			// name() => string
			"name": &tengo.UserFunction{
				Name:  "name",
				Value: stdlib.FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &tengo.UserFunction{
				Name:  "readdirnames",
				Value: stdlib.FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &tengo.UserFunction{
				Name:  "sync",
				Value: stdlib.FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &tengo.UserFunction{
				Name:  "write",
				Value: stdlib.FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &tengo.UserFunction{
				Name:  "write_string",
				Value: stdlib.FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &tengo.UserFunction{
				Name:  "read",
				Value: stdlib.FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &tengo.UserFunction{
				Name: "chmod",
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
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &tengo.UserFunction{
				Name: "seek",
				Value: func(args ...tengo.Object) (tengo.Object, error) {
					if len(args) != 2 {
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
					i2, ok := tengo.ToInt(args[1])
					if !ok {
						return nil, tengo.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &tengo.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &tengo.UserFunction{
				Name: "stat",
				Value: func(args ...tengo.Object) (tengo.Object, error) {
					if len(args) != 0 {
						return nil, tengo.ErrWrongNumArguments
					}
					return osStat(&tengo.String{Value: file.Name()})
				},
			},
		},
	}
}
