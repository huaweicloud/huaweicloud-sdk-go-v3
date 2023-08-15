package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDcPointRequest Request Object
type DeleteDcPointRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	// 采集点位表id，创建点位表时设置，数据源下唯一。
	PointId string `json:"point_id"`

	// 设备id
	DeviceId *string `json:"device_id,omitempty"`

	// 设备服务属性，允许中、数字、英文大小写、下划线、中划线
	Property *string `json:"property,omitempty"`
}

func (o DeleteDcPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDcPointRequest struct{}"
	}

	return strings.Join([]string{"DeleteDcPointRequest", string(data)}, " ")
}
