package model

import (
	"encoding/json"

	"strings"
)

type XCodeError struct {
	// 错误码
	Code *string `json:"code,omitempty"`
	// 错误信息
	Msg *string `json:"msg,omitempty"`
}

func (o XCodeError) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "XCodeError struct{}"
	}

	return strings.Join([]string{"XCodeError", string(data)}, " ")
}
