package model

import (
	"encoding/json"

	"strings"
)

type CompressResponse struct {
	// GZIP压缩开关。0关闭。1打开

	CompressSwitch int32 `json:"compress_switch"`

	CompressRules *CompressRules `json:"compress_rules,omitempty"`
}

func (o CompressResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompressResponse struct{}"
	}

	return strings.Join([]string{"CompressResponse", string(data)}, " ")
}
