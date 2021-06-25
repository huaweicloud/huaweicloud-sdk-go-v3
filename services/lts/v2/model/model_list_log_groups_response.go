package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListLogGroupsResponse struct {
	// 日志组信息。

	LogGroups      *[]LogGroup `json:"log_groups,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListLogGroupsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLogGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListLogGroupsResponse", string(data)}, " ")
}
