package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkDevicesRequest Request Object
type ListNetworkDevicesRequest struct {

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 根据网络设备状态查询，支持多值查询
	Status *[]NetworkDeviceStatus `json:"status,omitempty"`

	// 根据ID过滤，支持多值查询，查询条件格式：id=xxx&id=xxx。
	Id *[]string `json:"id,omitempty"`
}

func (o ListNetworkDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListNetworkDevicesRequest", string(data)}, " ")
}
