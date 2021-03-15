package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAsyncHistoryCommandsResponse struct {
	// 设备命令列表。
	Commands       *[]AsyncDeviceCommand `json:"commands,omitempty"`
	Page           *Page                 `json:"page,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAsyncHistoryCommandsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAsyncHistoryCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncHistoryCommandsResponse", string(data)}, " ")
}
