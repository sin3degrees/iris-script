package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/sin3degrees/tengo/v2"
	"github.com/sin3degrees/tengo/v2/stdlib"
)

var osModule = map[string]interface{}{
	"O_RDONLY":          int64(os.O_RDONLY),
	"O_WRONLY":          int64(os.O_WRONLY),
	"O_RDWR":            int64(os.O_RDWR),
	"O_APPEND":          int64(os.O_APPEND),
	"O_CREATE":          int64(os.O_CREATE),
	"O_EXCL":            int64(os.O_EXCL),
	"O_SYNC":            int64(os.O_SYNC),
	"O_TRUNC":           int64(os.O_TRUNC),
	"ModeDir":           int64(os.ModeDir),
	"ModeAppend":        int64(os.ModeAppend),
	"ModeExclusive":     int64(os.ModeExclusive),
	"ModeTemporary":     int64(os.ModeTemporary),
	"ModeSymlink":       int64(os.ModeSymlink),
	"ModeDevice":        int64(os.ModeDevice),
	"ModeNamedPipe":     int64(os.ModeNamedPipe),
	"ModeSocket":        int64(os.ModeSocket),
	"ModeSetuid":        int64(os.ModeSetuid),
	"ModeSetgid":        int64(os.ModeSetgid),
	"ModeCharDevice":    int64(os.ModeCharDevice),
	"ModeSticky":        int64(os.ModeSticky),
	"ModeType":          int64(os.ModeType),
	"ModePerm":          int64(os.ModePerm),
	"PathSeparator":     os.PathSeparator,
	"PathListSeparator": os.PathListSeparator,
	"DevNull":           os.DevNull,
	"SeekStart":         int64(io.SeekStart),
	"SeekCurrent":       int64(io.SeekCurrent),
	"SeekEnd":           int64(io.SeekEnd),
	"Args":              osArgs,
	"Chdir":             stdlib.FuncASRE(os.Chdir),
	"Chmod":             osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"Chown":             stdlib.FuncASIIRE(os.Chown),
	"Clearenv":          stdlib.FuncAR(os.Clearenv),
	"Environ":           stdlib.FuncARSs(os.Environ),
	"Exit":              stdlib.FuncAIR(os.Exit),
	"ExpandEnv":         osExpandEnv,
	"Getegid":           stdlib.FuncARI(os.Getegid),
	"Getenv":            stdlib.FuncASRS(os.Getenv),
	"Geteuid":           stdlib.FuncARI(os.Geteuid),
	"Getgid":            stdlib.FuncARI(os.Getgid),
	"Getgroups":         stdlib.FuncARIsE(os.Getgroups),
	"Getpagesize":       stdlib.FuncARI(os.Getpagesize),
	"Getpid":            stdlib.FuncARI(os.Getpid),
	"Getppid":           stdlib.FuncARI(os.Getppid),
	"Getuid":            stdlib.FuncARI(os.Getuid),
	"Getwd":             stdlib.FuncARSE(os.Getwd),
	"Hostname":          stdlib.FuncARSE(os.Hostname),
	"Lchown":            stdlib.FuncASIIRE(os.Lchown),
	"Link":              stdlib.FuncASSRE(os.Link),
	"LookupEnv":         osLookupEnv,
	"mkdir":             osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all":         osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"Readlink":          stdlib.FuncASRSE(os.Readlink),
	"Remove":            stdlib.FuncASRE(os.Remove),
	"RemoveAll":         stdlib.FuncASRE(os.RemoveAll),
	"Rename":            stdlib.FuncASSRE(os.Rename),
	"Setenv":            stdlib.FuncASSRE(os.Setenv),
	"Symlink":           stdlib.FuncASSRE(os.Symlink),
	"TempDir":           stdlib.FuncARS(os.TempDir),
	"Truncate":          stdlib.FuncASI64RE(os.Truncate),
	"Unsetenv":          stdlib.FuncASRE(os.Unsetenv),
	"Create":            osCreate,
	"Open":              osOpen,
	"OpenFile":          osOpenFile,
	"FindProcess":       osFindProcess,
	"StartProcess":      osStartProcess,
	"LookPath":          stdlib.FuncASRSE(exec.LookPath),
	"Exec":              osExec,
	"Stat":              osStat,
	"ReadFile":          osReadFile,
}

func osReadFile(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	fname, ok := tengo.ToString(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > tengo.MaxBytesLen {
		return nil, tengo.ErrBytesLimit
	}
	return &tengo.Bytes{Value: bytes}, nil
}

func osStat(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	fname, ok := tengo.ToString(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &tengo.ImmutableMap{
		Value: map[string]tengo.Object{
			"name":  &tengo.String{Value: stat.Name()},
			"mtime": &tengo.Time{Value: stat.ModTime()},
			"size":  &tengo.Int{Value: stat.Size()},
			"mode":  &tengo.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = tengo.TrueValue
	} else {
		fstat.Value["directory"] = tengo.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 3 {
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
	i3, ok := tengo.ToInt(args[2])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 0 {
		return nil, tengo.ErrWrongNumArguments
	}
	arr := &tengo.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > tengo.MaxStringLen {
			return nil, tengo.ErrStringLimit
		}
		arr.Value = append(arr.Value, &tengo.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *tengo.UserFunction {
	return &tengo.UserFunction{
		Name: name,
		Value: func(args ...tengo.Object) (tengo.Object, error) {
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
			i2, ok := tengo.ToInt64(args[1])
			if !ok {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return tengo.FalseValue, nil
	}
	if len(res) > tengo.MaxStringLen {
		return nil, tengo.ErrStringLimit
	}
	return &tengo.String{Value: res}, nil
}

func osExpandEnv(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
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
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > tengo.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > tengo.MaxStringLen {
		return nil, tengo.ErrStringLimit
	}
	return &tengo.String{Value: s}, nil
}

func osExec(args ...tengo.Object) (tengo.Object, error) {
	if len(args) == 0 {
		return nil, tengo.ErrWrongNumArguments
	}
	name, ok := tengo.ToString(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := tengo.ToString(arg)
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	i1, ok := tengo.ToInt(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 4 {
		return nil, tengo.ErrWrongNumArguments
	}
	name, ok := tengo.ToString(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *tengo.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *tengo.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := tengo.ToString(args[2])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *tengo.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *tengo.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []tengo.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*tengo.String)
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
