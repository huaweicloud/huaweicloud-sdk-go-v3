package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsForOrganizationUnitRequest Request Object
type ListControlsForOrganizationUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationUnitId string `json:"managed_organization_unit_id"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListControlsForOrganizationUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsForOrganizationUnitRequest struct{}"
	}

	return strings.Join([]string{"ListControlsForOrganizationUnitRequest", string(data)}, " ")
}
