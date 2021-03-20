package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApiVersionsResponse struct {
	// 描述version相关对象的列表

	Versions       *[]VersionDetail `json:"versions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
