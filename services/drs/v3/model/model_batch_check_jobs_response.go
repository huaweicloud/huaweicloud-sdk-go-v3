package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCheckJobsResponse struct {
	// 预检查响应体

	Results *[]PostPreCheckResp `json:"results,omitempty"`
	// 总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCheckJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCheckJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckJobsResponse", string(data)}, " ")
}
