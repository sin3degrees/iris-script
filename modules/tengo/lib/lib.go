package lib

import "github.com/sin3degrees/tengo/v2"

func Load(s *tengo.Script) {
	// 标准库 stdlib
	s.Add("base64", base64Module)
	s.Add("fmt", fmtModule)
	s.Add("hex", hexModule)
	s.Add("json", jsonModule)
	s.Add("math", mathModule)
	s.Add("os", osModule)
	s.Add("rand", randModule)
	s.Add("text", textModule)
	s.Add("times", timesModule)
	// 自定义库 selflib
	s.Add("orm", ormModule)
}
