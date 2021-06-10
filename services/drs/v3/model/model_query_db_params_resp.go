package model

import (
	"encoding/json"

	"strings"
)

// 数据库参数信息响应体
type QueryDbParamsResp struct {
	Params *[]Params `json:"params,omitempty"`
}

func (o QueryDbParamsResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryDbParamsResp struct{}"
	}

	return strings.Join([]string{"QueryDbParamsResp", string(data)}, " ")
}
