package lib

import (
	"encoding/base64"

	"github.com/sin3degrees/tengo/v2/stdlib"
)

var base64Module = map[string]interface{}{
	"Encode":       stdlib.FuncAYRS(base64.StdEncoding.EncodeToString),
	"Decode":       stdlib.FuncASRYE(base64.StdEncoding.DecodeString),
	"RawEncode":    stdlib.FuncAYRS(base64.RawStdEncoding.EncodeToString),
	"RawDecode":    stdlib.FuncASRYE(base64.RawStdEncoding.DecodeString),
	"UrlEncode":    stdlib.FuncAYRS(base64.URLEncoding.EncodeToString),
	"UrlDecode":    stdlib.FuncASRYE(base64.URLEncoding.DecodeString),
	"RawUrlEncode": stdlib.FuncAYRS(base64.RawURLEncoding.EncodeToString),
	"RawUrlDecode": stdlib.FuncASRYE(base64.RawURLEncoding.DecodeString),
}
