package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPointsRequest Request Object
type ShowPointsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`
}

func (o ShowPointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPointsRequest struct{}"
	}

	return strings.Join([]string{"ShowPointsRequest", string(data)}, " ")
}
