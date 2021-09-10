package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowJobsResponse struct {
	// 作业数

	Total *int32 `json:"total,omitempty"`
	// 作业列表，请参见jobs参数说明

	Jobs *[]Job `json:"jobs,omitempty"`
	// 当为“true”时返回精简消息，即作业参数只返回参数名和值，不返回参数的“size”、“type”、“id”等属性

	Simple *bool `json:"simple,omitempty"`
	// 返回指定页号的作业

	PageNo *int32 `json:"page_no,omitempty"`
	// 每页作业数

	PageSize       *int32 `json:"page_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobsResponse struct{}"
	}

	return strings.Join([]string{"ShowJobsResponse", string(data)}, " ")
}
