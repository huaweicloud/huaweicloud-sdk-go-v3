package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchResetPasswordResponse struct {
	// 总数

	Count *int32 `json:"count,omitempty"`
	// 批量修改任务返回列表

	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchResetPasswordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchResetPasswordResponse struct{}"
	}

	return strings.Join([]string{"BatchResetPasswordResponse", string(data)}, " ")
}
