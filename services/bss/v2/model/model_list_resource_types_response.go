package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourceTypesResponse struct {
	// 资源类型信息，具体参见表3。
	ResourceTypes  *[]ResourceType `json:"resource_types,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListResourceTypesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourceTypesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceTypesResponse", string(data)}, " ")
}
