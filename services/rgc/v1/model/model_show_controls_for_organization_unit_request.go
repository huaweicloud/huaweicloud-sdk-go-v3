package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlsForOrganizationUnitRequest Request Object
type ShowControlsForOrganizationUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationUnitId string `json:"managed_organization_unit_id"`

	// 控制策略ID。
	ControlId string `json:"control_id"`
}

func (o ShowControlsForOrganizationUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlsForOrganizationUnitRequest struct{}"
	}

	return strings.Join([]string{"ShowControlsForOrganizationUnitRequest", string(data)}, " ")
}
