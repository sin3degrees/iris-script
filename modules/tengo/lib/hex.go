package lib

import (
	"encoding/hex"

	"github.com/sin3degrees/tengo/v2/stdlib"
)

var hexModule = map[string]interface{}{
	"Encode": stdlib.FuncAYRS(hex.EncodeToString),
	"Decode": stdlib.FuncASRYE(hex.DecodeString),
}
