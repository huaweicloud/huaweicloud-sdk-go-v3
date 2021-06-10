package model

import (
	"encoding/json"

	"strings"
)

type CompressRequest struct {
	// GZIP压缩开关。0关闭。1打开

	CompressSwitch *int32 `json:"compress_switch,omitempty"`
}

func (o CompressRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompressRequest struct{}"
	}

	return strings.Join([]string{"CompressRequest", string(data)}, " ")
}
