package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShipperAuthorizationsRequest Request Object
type ListShipperAuthorizationsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据源
	SourceRegion *string `json:"source_region,omitempty"`

	// 目的投递类型
	DestinationShipperType *string `json:"destination_shipper_type,omitempty"`

	// 数据投递状态
	ShipperStatus *string `json:"shipper_status,omitempty"`

	// 授权状态
	AuthStatus *string `json:"auth_status,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListShipperAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShipperAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"ListShipperAuthorizationsRequest", string(data)}, " ")
}
