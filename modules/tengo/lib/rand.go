package lib

import (
	"math/rand"

	"github.com/sin3degrees/tengo/v2"
	"github.com/sin3degrees/tengo/v2/stdlib"
)

var randModule = map[string]interface{}{
	"Int":       stdlib.FuncARI64(rand.Int63),
	"Float":     stdlib.FuncARF(rand.Float64),
	"Intn":      stdlib.FuncAI64RI64(rand.Int63n),
	"ExpFloat":  stdlib.FuncARF(rand.ExpFloat64),
	"NormFloat": stdlib.FuncARF(rand.NormFloat64),
	"Perm":      stdlib.FuncAIRIs(rand.Perm),
	"Seed":      stdlib.FuncAI64R(rand.Seed),
	"Read": func(args ...tengo.Object) (ret tengo.Object, err error) {
		if len(args) != 1 {
			return nil, tengo.ErrWrongNumArguments
		}
		y1, ok := args[0].(*tengo.Bytes)
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "bytes",
				Found:    args[0].TypeName(),
			}
		}
		res, err := rand.Read(y1.Value)
		if err != nil {
			ret = wrapError(err)
			return
		}
		return &tengo.Int{Value: int64(res)}, nil
	},
	"Rand": func(args ...tengo.Object) (tengo.Object, error) {
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
		src := rand.NewSource(i1)
		return randRand(rand.New(src)), nil
	},
}

func randRand(r *rand.Rand) *tengo.ImmutableMap {
	return &tengo.ImmutableMap{
		Value: map[string]tengo.Object{
			"int": &tengo.UserFunction{
				Name:  "int",
				Value: stdlib.FuncARI64(r.Int63),
			},
			"float": &tengo.UserFunction{
				Name:  "float",
				Value: stdlib.FuncARF(r.Float64),
			},
			"intn": &tengo.UserFunction{
				Name:  "intn",
				Value: stdlib.FuncAI64RI64(r.Int63n),
			},
			"exp_float": &tengo.UserFunction{
				Name:  "exp_float",
				Value: stdlib.FuncARF(r.ExpFloat64),
			},
			"norm_float": &tengo.UserFunction{
				Name:  "norm_float",
				Value: stdlib.FuncARF(r.NormFloat64),
			},
			"perm": &tengo.UserFunction{
				Name:  "perm",
				Value: stdlib.FuncAIRIs(r.Perm),
			},
			"seed": &tengo.UserFunction{
				Name:  "seed",
				Value: stdlib.FuncAI64R(r.Seed),
			},
			"read": &tengo.UserFunction{
				Name: "read",
				Value: func(args ...tengo.Object) (
					ret tengo.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, tengo.ErrWrongNumArguments
					}
					y1, ok := args[0].(*tengo.Bytes)
					if !ok {
						return nil, tengo.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &tengo.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
