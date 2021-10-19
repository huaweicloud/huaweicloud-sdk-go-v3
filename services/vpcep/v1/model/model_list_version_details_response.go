package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListVersionDetailsResponse struct {
	// VPC终端节点版本信息列表。

	Versions       *[]Versions `json:"versions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListVersionDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVersionDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListVersionDetailsResponse", string(data)}, " ")
}
