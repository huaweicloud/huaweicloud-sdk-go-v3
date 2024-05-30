package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsForOrganizationalUnitRequest Request Object
type ListControlsForOrganizationalUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationalUnitId string `json:"managed_organizational_unit_id"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListControlsForOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsForOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"ListControlsForOrganizationalUnitRequest", string(data)}, " ")
}
