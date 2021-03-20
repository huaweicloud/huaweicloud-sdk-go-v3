package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAsyncCommandsResponse struct {
	// 设备命令列表。

	Commands *[]AsyncDeviceCommand `json:"commands,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAsyncCommandsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAsyncCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncCommandsResponse", string(data)}, " ")
}
