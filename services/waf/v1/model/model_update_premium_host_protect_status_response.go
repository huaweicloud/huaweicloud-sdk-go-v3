package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePremiumHostProtectStatusResponse struct {
	// 防护状态

	ProtectStatus  *int32 `json:"protect_status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePremiumHostProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostProtectStatusResponse", string(data)}, " ")
}
