package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchChangeDataResponse struct {
	// 批量数据加工响应列表

	Results *[]DataTransformationResp `json:"results,omitempty"`
	// 总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchChangeDataResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchChangeDataResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeDataResponse", string(data)}, " ")
}
