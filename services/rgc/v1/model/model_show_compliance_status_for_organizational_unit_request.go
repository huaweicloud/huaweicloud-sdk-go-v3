package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceStatusForOrganizationalUnitRequest Request Object
type ShowComplianceStatusForOrganizationalUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationalUnitId string `json:"managed_organizational_unit_id"`

	// 控制策略ID。
	ControlId *string `json:"control_id,omitempty"`
}

func (o ShowComplianceStatusForOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceStatusForOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"ShowComplianceStatusForOrganizationalUnitRequest", string(data)}, " ")
}
