package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceStatusForOrganizationUnitResponse Response Object
type ShowComplianceStatusForOrganizationUnitResponse struct {

	// 合规状态。
	ComplianceStatus *string `json:"compliance_status,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowComplianceStatusForOrganizationUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceStatusForOrganizationUnitResponse struct{}"
	}

	return strings.Join([]string{"ShowComplianceStatusForOrganizationUnitResponse", string(data)}, " ")
}
