package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateMappingRequest struct {
	// 待更新的映射ID。

	Id string `json:"id"`

	Body *KeystoneUpdateMappingRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateMappingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateMappingRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateMappingRequest", string(data)}, " ")
}
