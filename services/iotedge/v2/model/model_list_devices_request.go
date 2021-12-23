package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDevicesRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 父设备ID,对应之前的gatewayId的概念，传该参数时代表查询网关下的子设备，不传代表节点下的

	GatewayId *string `json:"gateway_id,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数，默认值为10，取值区间为1-1000

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListDevicesRequest", string(data)}, " ")
}
