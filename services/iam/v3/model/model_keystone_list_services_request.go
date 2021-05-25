package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListServicesRequest struct {
	// 服务类型。

	Type *string `json:"type,omitempty"`
}

func (o KeystoneListServicesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListServicesRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListServicesRequest", string(data)}, " ")
}
