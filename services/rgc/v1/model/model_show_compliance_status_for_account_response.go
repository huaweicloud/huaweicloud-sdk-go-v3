package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceStatusForAccountResponse Response Object
type ShowComplianceStatusForAccountResponse struct {

	// 合规状态。
	ComplianceStatus *string `json:"compliance_status,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowComplianceStatusForAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceStatusForAccountResponse struct{}"
	}

	return strings.Join([]string{"ShowComplianceStatusForAccountResponse", string(data)}, " ")
}
