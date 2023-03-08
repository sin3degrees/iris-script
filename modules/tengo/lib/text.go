package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/sin3degrees/tengo/v2"
	"github.com/sin3degrees/tengo/v2/stdlib"
)

var textModule = map[string]interface{}{
	"REMatch":      textREMatch,
	"REFind":       textREFind,
	"REReplace":    textREReplace,
	"RESplit":      textRESplit,
	"RECompile":    textRECompile,
	"Compare":      stdlib.FuncASSRI(strings.Compare),
	"Contains":     stdlib.FuncASSRB(strings.Contains),
	"ContainsAny":  stdlib.FuncASSRB(strings.ContainsAny),
	"Count":        stdlib.FuncASSRI(strings.Count),
	"EqualFold":    stdlib.FuncASSRB(strings.EqualFold),
	"fields":       stdlib.FuncASRSs(strings.Fields),
	"HasPrefix":    stdlib.FuncASSRB(strings.HasPrefix),
	"HasSuffix":    stdlib.FuncASSRB(strings.HasSuffix),
	"Index":        stdlib.FuncASSRI(strings.Index),
	"IndexAny":     stdlib.FuncASSRI(strings.IndexAny),
	"Join":         textJoin,
	"LastIndex":    stdlib.FuncASSRI(strings.LastIndex),
	"LastIndexAny": stdlib.FuncASSRI(strings.LastIndexAny),
	"Repeat":       textRepeat,
	"Replace":      textReplace,
	"Substr":       textSubstring,
	"Split":        stdlib.FuncASSRSs(strings.Split),
	"SplitAfter":   stdlib.FuncASSRSs(strings.SplitAfter),
	"SplitAfterN":  stdlib.FuncASSIRSs(strings.SplitAfterN),
	"SplitN":       stdlib.FuncASSIRSs(strings.SplitN),
	"Title":        stdlib.FuncASRS(strings.Title),
	"ToLower":      stdlib.FuncASRS(strings.ToLower),
	"ToTitle":      stdlib.FuncASRS(strings.ToTitle),
	"ToUpper":      stdlib.FuncASRS(strings.ToUpper),
	"PadLeft":      textPadLeft,
	"PadRight":     textPadRight,
	"Trim":         stdlib.FuncASSRS(strings.Trim),
	"TrimLeft":     stdlib.FuncASSRS(strings.TrimLeft),
	"TrimPrefix":   stdlib.FuncASSRS(strings.TrimPrefix),
	"TrimRight":    stdlib.FuncASSRS(strings.TrimRight),
	"TrimSpace":    stdlib.FuncASRS(strings.TrimSpace),
	"TrimSuffix":   stdlib.FuncASSRS(strings.TrimSuffix),
	"Atoi":         stdlib.FuncASRIE(strconv.Atoi),
	"FormatBool":   textFormatBool,
	"FormatFloat":  textFormatFloat,
	"FormatInt":    textFormatInt,
	"Itoa":         stdlib.FuncAIRS(strconv.Itoa),
	"ParseBool":    textParseBool,
	"ParseFloat":   textParseFloat,
	"ParseInt":     textParseInt,
	"Quote":        stdlib.FuncASRS(strconv.Quote),
	"Unquote":      stdlib.FuncASRSE(strconv.Unquote),
}

func textREMatch(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	matched, err := regexp.MatchString(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	if matched {
		ret = tengo.TrueValue
	} else {
		ret = tengo.FalseValue
	}

	return
}

func textREFind(args ...tengo.Object) (ret tengo.Object, err error) {
	numArgs := len(args)
	if numArgs != 2 && numArgs != 3 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if numArgs < 3 {
		m := re.FindStringSubmatchIndex(s2)
		if m == nil {
			ret = tengo.UndefinedValue
			return
		}

		arr := &tengo.Array{}
		for i := 0; i < len(m); i += 2 {
			arr.Value = append(arr.Value,
				&tengo.ImmutableMap{Value: map[string]tengo.Object{
					"text":  &tengo.String{Value: s2[m[i]:m[i+1]]},
					"begin": &tengo.Int{Value: int64(m[i])},
					"end":   &tengo.Int{Value: int64(m[i+1])},
				}})
		}

		ret = &tengo.Array{Value: []tengo.Object{arr}}

		return
	}

	i3, ok := tengo.ToInt(args[2])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	m := re.FindAllStringSubmatchIndex(s2, i3)
	if m == nil {
		ret = tengo.UndefinedValue
		return
	}

	arr := &tengo.Array{}
	for _, m := range m {
		subMatch := &tengo.Array{}
		for i := 0; i < len(m); i += 2 {
			subMatch.Value = append(subMatch.Value,
				&tengo.ImmutableMap{Value: map[string]tengo.Object{
					"text":  &tengo.String{Value: s2[m[i]:m[i+1]]},
					"begin": &tengo.Int{Value: int64(m[i])},
					"end":   &tengo.Int{Value: int64(m[i+1])},
				}})
		}

		arr.Value = append(arr.Value, subMatch)
	}

	ret = arr

	return
}

func textREReplace(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 3 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s3, ok := tengo.ToString(args[2])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
	} else {
		s, ok := doTextRegexpReplace(re, s2, s3)
		if !ok {
			return nil, tengo.ErrStringLimit
		}

		ret = &tengo.String{Value: s}
	}

	return
}

func textRESplit(args ...tengo.Object) (ret tengo.Object, err error) {
	numArgs := len(args)
	if numArgs != 2 && numArgs != 3 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	var i3 = -1
	if numArgs > 2 {
		i3, ok = tengo.ToInt(args[2])
		if !ok {
			err = tengo.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	arr := &tengo.Array{}
	for _, s := range re.Split(s2, i3) {
		arr.Value = append(arr.Value, &tengo.String{Value: s})
	}

	ret = arr

	return
}

func textRECompile(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
	} else {
		ret = makeTextRegexp(re)
	}

	return
}

func textReplace(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 4 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s3, ok := tengo.ToString(args[2])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := tengo.ToInt(args[3])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	s, ok := doTextReplace(s1, s2, s3, i4)
	if !ok {
		err = tengo.ErrStringLimit
		return
	}

	ret = &tengo.String{Value: s}

	return
}

func textSubstring(args ...tengo.Object) (ret tengo.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	strlen := len(s1)
	i3 := strlen
	if argslen == 3 {
		i3, ok = tengo.ToInt(args[2])
		if !ok {
			err = tengo.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	if i2 > i3 {
		err = tengo.ErrInvalidIndexType
		return
	}

	if i2 < 0 {
		i2 = 0
	} else if i2 > strlen {
		i2 = strlen
	}

	if i3 < 0 {
		i3 = 0
	} else if i3 > strlen {
		i3 = strlen
	}

	ret = &tengo.String{Value: s1[i2:i3]}

	return
}

func textPadLeft(args ...tengo.Object) (ret tengo.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if i2 > tengo.MaxStringLen {
		return nil, tengo.ErrStringLimit
	}

	sLen := len(s1)
	if sLen >= i2 {
		ret = &tengo.String{Value: s1}
		return
	}

	s3 := " "
	if argslen == 3 {
		s3, ok = tengo.ToString(args[2])
		if !ok {
			err = tengo.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "string(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	padStrLen := len(s3)
	if padStrLen == 0 {
		ret = &tengo.String{Value: s1}
		return
	}

	padCount := ((i2 - padStrLen) / padStrLen) + 1
	retStr := strings.Repeat(s3, padCount) + s1
	ret = &tengo.String{Value: retStr[len(retStr)-i2:]}

	return
}

func textPadRight(args ...tengo.Object) (ret tengo.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if i2 > tengo.MaxStringLen {
		return nil, tengo.ErrStringLimit
	}

	sLen := len(s1)
	if sLen >= i2 {
		ret = &tengo.String{Value: s1}
		return
	}

	s3 := " "
	if argslen == 3 {
		s3, ok = tengo.ToString(args[2])
		if !ok {
			err = tengo.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "string(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	padStrLen := len(s3)
	if padStrLen == 0 {
		ret = &tengo.String{Value: s1}
		return
	}

	padCount := ((i2 - padStrLen) / padStrLen) + 1
	retStr := s1 + strings.Repeat(s3, padCount)
	ret = &tengo.String{Value: retStr[:i2]}

	return
}

func textRepeat(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		return nil, tengo.ErrWrongNumArguments
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
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

	if len(s1)*i2 > tengo.MaxStringLen {
		return nil, tengo.ErrStringLimit
	}

	return &tengo.String{Value: strings.Repeat(s1, i2)}, nil
}

func textJoin(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		return nil, tengo.ErrWrongNumArguments
	}

	var slen int
	var ss1 []string
	switch arg0 := args[0].(type) {
	case *tengo.Array:
		for idx, a := range arg0.Value {
			as, ok := tengo.ToString(a)
			if !ok {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			slen += len(as)
			ss1 = append(ss1, as)
		}
	case *tengo.ImmutableArray:
		for idx, a := range arg0.Value {
			as, ok := tengo.ToString(a)
			if !ok {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			slen += len(as)
			ss1 = append(ss1, as)
		}
	default:
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	// make sure output length does not exceed the limit
	if slen+len(s2)*(len(ss1)-1) > tengo.MaxStringLen {
		return nil, tengo.ErrStringLimit
	}

	return &tengo.String{Value: strings.Join(ss1, s2)}, nil
}

func textFormatBool(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	b1, ok := args[0].(*tengo.Bool)
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bool",
			Found:    args[0].TypeName(),
		}
		return
	}

	if b1 == tengo.TrueValue {
		ret = &tengo.String{Value: "true"}
	} else {
		ret = &tengo.String{Value: "false"}
	}

	return
}

func textFormatFloat(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 4 {
		err = tengo.ErrWrongNumArguments
		return
	}

	f1, ok := args[0].(*tengo.Float)
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "float",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := tengo.ToInt(args[2])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := tengo.ToInt(args[3])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &tengo.String{Value: strconv.FormatFloat(f1.Value, s2[0], i3, i4)}

	return
}

func textFormatInt(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].(*tengo.Int)
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tengo.String{Value: strconv.FormatInt(i1.Value, i2)}

	return
}

func textParseBool(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].(*tengo.String)
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
		return
	}

	parsed, err := strconv.ParseBool(s1.Value)
	if err != nil {
		ret = wrapError(err)
		return
	}

	if parsed {
		ret = tengo.TrueValue
	} else {
		ret = tengo.FalseValue
	}

	return
}

func textParseFloat(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].(*tengo.String)
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := strconv.ParseFloat(s1.Value, i2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &tengo.Float{Value: parsed}

	return
}

func textParseInt(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 3 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].(*tengo.String)
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := tengo.ToInt(args[2])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	parsed, err := strconv.ParseInt(s1.Value, i2, i3)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &tengo.Int{Value: parsed}

	return
}

// Modified implementation of strings.Replace
// to limit the maximum length of output string.
func doTextReplace(s, old, new string, n int) (string, bool) {
	if old == new || n == 0 {
		return s, true // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s, true // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}

		ssj := s[start:j]
		if w+len(ssj)+len(new) > tengo.MaxStringLen {
			return "", false
		}

		w += copy(t[w:], ssj)
		w += copy(t[w:], new)
		start = j + len(old)
	}

	ss := s[start:]
	if w+len(ss) > tengo.MaxStringLen {
		return "", false
	}

	w += copy(t[w:], ss)

	return string(t[0:w]), true
}
