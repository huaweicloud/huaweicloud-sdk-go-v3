package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchListJobDetailsResponse struct {
	// 任务数量

	Count *int32 `json:"count,omitempty"`
	// 任务详细信息

	Results        *[]QueryJobResp `json:"results,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchListJobDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchListJobDetailsResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobDetailsResponse", string(data)}, " ")
}
