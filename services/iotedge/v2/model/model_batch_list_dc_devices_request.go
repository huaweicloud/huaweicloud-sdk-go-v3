package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListDcDevicesRequest Request Object
type BatchListDcDevicesRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	// 设备标识码。
	DeviceId *string `json:"device_id,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认值为10，取值区间为1-1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o BatchListDcDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListDcDevicesRequest struct{}"
	}

	return strings.Join([]string{"BatchListDcDevicesRequest", string(data)}, " ")
}
