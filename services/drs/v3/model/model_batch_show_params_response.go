package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchShowParamsResponse struct {
	// 查询数据库参数响应体

	ParamsList *[]QueryDbParamsResp `json:"params_list,omitempty"`
	// 总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchShowParamsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchShowParamsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowParamsResponse", string(data)}, " ")
}
