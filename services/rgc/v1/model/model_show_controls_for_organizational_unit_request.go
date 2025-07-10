package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlsForOrganizationalUnitRequest Request Object
type ShowControlsForOrganizationalUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationalUnitId string `json:"managed_organizational_unit_id"`

	// 控制策略ID。
	ControlId string `json:"control_id"`
}

func (o ShowControlsForOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlsForOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"ShowControlsForOrganizationalUnitRequest", string(data)}, " ")
}
