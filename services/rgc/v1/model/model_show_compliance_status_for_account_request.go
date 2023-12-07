package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceStatusForAccountRequest Request Object
type ShowComplianceStatusForAccountRequest struct {

	// 账号ID。
	ManagedAccountId string `json:"managed_account_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ShowComplianceStatusForAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceStatusForAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowComplianceStatusForAccountRequest", string(data)}, " ")
}
