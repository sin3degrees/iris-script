package lib

import (
	"math"

	"github.com/sin3degrees/tengo/v2/stdlib"
)

var mathModule = map[string]interface{}{
	"E":         math.E,
	"Pi":        math.Pi,
	"Phi":       math.Phi,
	"Sqrt2":     math.Sqrt2,
	"SqrtE":     math.SqrtE,
	"SqrtPi":    math.SqrtPi,
	"SqrtPhi":   math.SqrtPhi,
	"Ln2":       math.Ln2,
	"Ln10":      math.Ln10,
	"Log2E":     math.Log2E,
	"Log10E":    math.Log10E,
	"Abs":       stdlib.FuncAFRF(math.Abs),
	"Acos":      stdlib.FuncAFRF(math.Acos),
	"Acosh":     stdlib.FuncAFRF(math.Acosh),
	"Asin":      stdlib.FuncAFRF(math.Asin),
	"Asinh":     stdlib.FuncAFRF(math.Asinh),
	"Atan":      stdlib.FuncAFRF(math.Atan),
	"Atan2":     stdlib.FuncAFFRF(math.Atan2),
	"Atanh":     stdlib.FuncAFRF(math.Atanh),
	"Cbrt":      stdlib.FuncAFRF(math.Cbrt),
	"Ceil":      stdlib.FuncAFRF(math.Ceil),
	"Copysign":  stdlib.FuncAFFRF(math.Copysign),
	"Cos":       stdlib.FuncAFRF(math.Cos),
	"Cosh":      stdlib.FuncAFRF(math.Cosh),
	"Dim":       stdlib.FuncAFFRF(math.Dim),
	"Erf":       stdlib.FuncAFRF(math.Erf),
	"Erfc":      stdlib.FuncAFRF(math.Erfc),
	"Exp":       stdlib.FuncAFRF(math.Exp),
	"Exp2":      stdlib.FuncAFRF(math.Exp2),
	"Expm1":     stdlib.FuncAFRF(math.Expm1),
	"Floor":     stdlib.FuncAFRF(math.Floor),
	"Gamma":     stdlib.FuncAFRF(math.Gamma),
	"Hypot":     stdlib.FuncAFFRF(math.Hypot),
	"Ilogb":     stdlib.FuncAFRI(math.Ilogb),
	"Inf":       stdlib.FuncAIRF(math.Inf),
	"IsInf":     stdlib.FuncAFIRB(math.IsInf),
	"IsNaN":     stdlib.FuncAFRB(math.IsNaN),
	"J0":        stdlib.FuncAFRF(math.J0),
	"J1":        stdlib.FuncAFRF(math.J1),
	"Jn":        stdlib.FuncAIFRF(math.Jn),
	"Ldexp":     stdlib.FuncAFIRF(math.Ldexp),
	"Log":       stdlib.FuncAFRF(math.Log),
	"Log10":     stdlib.FuncAFRF(math.Log10),
	"Log1p":     stdlib.FuncAFRF(math.Log1p),
	"Log2":      stdlib.FuncAFRF(math.Log2),
	"Logb":      stdlib.FuncAFRF(math.Logb),
	"Max":       stdlib.FuncAFFRF(math.Max),
	"Min":       stdlib.FuncAFFRF(math.Min),
	"Mod":       stdlib.FuncAFFRF(math.Mod),
	"NaN":       stdlib.FuncARF(math.NaN),
	"Nextafter": stdlib.FuncAFFRF(math.Nextafter),
	"Pow":       stdlib.FuncAFFRF(math.Pow),
	"Remainder": stdlib.FuncAFFRF(math.Remainder),
	"Signbit":   stdlib.FuncAFRB(math.Signbit),
	"Sin":       stdlib.FuncAFRF(math.Sin),
	"Sinh":      stdlib.FuncAFRF(math.Sinh),
	"Sqrt":      stdlib.FuncAFRF(math.Sqrt),
	"Tan":       stdlib.FuncAFRF(math.Tan),
	"Tanh":      stdlib.FuncAFRF(math.Tanh),
	"Trunc":     stdlib.FuncAFRF(math.Trunc),
	"Y0":        stdlib.FuncAFRF(math.Y0),
	"Y1":        stdlib.FuncAFRF(math.Y1),
	"Yn":        stdlib.FuncAIFRF(math.Yn),
}
