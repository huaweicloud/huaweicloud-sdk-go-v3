package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMetadataResponse struct {
	// 导入结果信息。

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMetadataResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMetadataResponse struct{}"
	}

	return strings.Join([]string{"CreateMetadataResponse", string(data)}, " ")
}
