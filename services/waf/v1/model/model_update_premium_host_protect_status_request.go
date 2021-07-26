package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePremiumHostProtectStatusRequest struct {
	// 独享模式域名ID

	HostId string `json:"host_id"`

	Body *UpdatePremiumHostProtectStatusRequestBody `json:"body,omitempty"`
}

func (o UpdatePremiumHostProtectStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostProtectStatusRequest", string(data)}, " ")
}
