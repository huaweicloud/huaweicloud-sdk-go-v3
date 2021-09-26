package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProcessesResponse struct {
	// 会话列表

	Processes *[]Process `json:"processes,omitempty"`
	// 总记录数

	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProcessesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProcessesResponse struct{}"
	}

	return strings.Join([]string{"ListProcessesResponse", string(data)}, " ")
}