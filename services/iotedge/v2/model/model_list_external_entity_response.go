package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListExternalEntityResponse struct {
	// 外部实体列表

	Externals      *[]ExternalEntityRespDto `json:"externals,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListExternalEntityResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListExternalEntityResponse struct{}"
	}

	return strings.Join([]string{"ListExternalEntityResponse", string(data)}, " ")
}
