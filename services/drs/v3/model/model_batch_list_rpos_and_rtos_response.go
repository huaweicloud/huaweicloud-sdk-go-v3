package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchListRposAndRtosResponse struct {
	// 总数

	Count *int32 `json:"count,omitempty"`
	// 批量查询RPO&RTO的响应体集合

	Results        *[]QueryRpoAndRtoResp `json:"results,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchListRposAndRtosResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchListRposAndRtosResponse struct{}"
	}

	return strings.Join([]string{"BatchListRposAndRtosResponse", string(data)}, " ")
}
