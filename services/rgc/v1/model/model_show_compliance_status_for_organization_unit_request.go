package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceStatusForOrganizationUnitRequest Request Object
type ShowComplianceStatusForOrganizationUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationUnitId string `json:"managed_organization_unit_id"`

	// 启用的控制策略信息。
	ControlId *string `json:"control_id,omitempty"`
}

func (o ShowComplianceStatusForOrganizationUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceStatusForOrganizationUnitRequest struct{}"
	}

	return strings.Join([]string{"ShowComplianceStatusForOrganizationUnitRequest", string(data)}, " ")
}
