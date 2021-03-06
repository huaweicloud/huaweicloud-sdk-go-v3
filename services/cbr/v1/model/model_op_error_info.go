package model

import (
	"encoding/json"

	"strings"
)

type OpErrorInfo struct {
	// 请参见[错误码](ErrorCode.xml)。

	Code *string `json:"code,omitempty"`
	// 错误信息

	Message *string `json:"message,omitempty"`
}

func (o OpErrorInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpErrorInfo struct{}"
	}

	return strings.Join([]string{"OpErrorInfo", string(data)}, " ")
}
