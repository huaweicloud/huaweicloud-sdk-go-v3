package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDcPointRequest Request Object
type UpdateDcPointRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	// 采集点位表id，创建点位表时设置，数据源下唯一。
	PointId string `json:"point_id"`

	Body *UpdateDcPointReqDto `json:"body,omitempty"`
}

func (o UpdateDcPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcPointRequest struct{}"
	}

	return strings.Join([]string{"UpdateDcPointRequest", string(data)}, " ")
}
