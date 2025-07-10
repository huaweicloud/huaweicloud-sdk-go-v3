package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceStatusForOrganizationalUnitResponse Response Object
type ShowComplianceStatusForOrganizationalUnitResponse struct {

	// 合规状态。
	ComplianceStatus *string `json:"compliance_status,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowComplianceStatusForOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceStatusForOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"ShowComplianceStatusForOrganizationalUnitResponse", string(data)}, " ")
}
