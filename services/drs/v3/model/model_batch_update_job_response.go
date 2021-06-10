package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchUpdateJobResponse struct {
	// 总数

	Count *int32 `json:"count,omitempty"`
	// 批量修改任务返回列表

	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchUpdateJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUpdateJobResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateJobResponse", string(data)}, " ")
}
