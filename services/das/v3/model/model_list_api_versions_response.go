package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApiVersionsResponse struct {
	// API版本详细信息列表。

	Versions       *[]ApiVersion `json:"versions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
