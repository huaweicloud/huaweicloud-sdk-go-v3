package model

import (
	"encoding/json"

	"strings"
)

// 迁移任务查询接口中返回的任务失败信息提示
type ErrorReasonResp struct {
	// 迁移失败的错误码。

	ErrorCode *string `json:"error_code,omitempty"`
	// 迁移失败的原因。

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ErrorReasonResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ErrorReasonResp struct{}"
	}

	return strings.Join([]string{"ErrorReasonResp", string(data)}, " ")
}
