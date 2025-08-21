package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceStatusForAccountRequest Request Object
type ShowComplianceStatusForAccountRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`

	// 控制策略ID。
	ControlId *string `json:"control_id,omitempty"`
}

func (o ShowComplianceStatusForAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceStatusForAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowComplianceStatusForAccountRequest", string(data)}, " ")
}
