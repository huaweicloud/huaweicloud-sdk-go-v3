package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 防护状态
type UpdatePremiumHostProtectStatusRequestBody struct {
	// 防护状态

	ProtectStatus *int32 `json:"protect_status,omitempty"`
}

func (o UpdatePremiumHostProtectStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostProtectStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostProtectStatusRequestBody", string(data)}, " ")
}
