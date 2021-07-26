package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePremiumHostProtectStatusResponse struct {
	// 防护状态

	ProtectStatus  *int32 `json:"protect_status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePremiumHostProtectStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostProtectStatusResponse", string(data)}, " ")
}
