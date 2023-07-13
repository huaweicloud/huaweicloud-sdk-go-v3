package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListDcPointsRequest Request Object
type BatchListDcPointsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	// 采集点位表id，创建点位表时设置，数据源下唯一。
	PointId *string `json:"point_id,omitempty"`

	// 点位名称，允许中、数字、英文大小写、下划线、中划线、#%()*特殊字符.模糊查询
	Name *string `json:"name,omitempty"`

	// 属性，允许中、数字、英文大小写、下划线、中划线，精确查询
	Property *string `json:"property,omitempty"`

	// 设备标识，精确查询
	DeviceId *string `json:"device_id,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认值为10，取值区间为1-1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o BatchListDcPointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListDcPointsRequest struct{}"
	}

	return strings.Join([]string{"BatchListDcPointsRequest", string(data)}, " ")
}
