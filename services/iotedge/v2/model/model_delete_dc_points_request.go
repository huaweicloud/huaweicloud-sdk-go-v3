package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDcPointsRequest Request Object
type DeleteDcPointsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	Body *DeleteDcPointsReqDto `json:"body,omitempty"`
}

func (o DeleteDcPointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDcPointsRequest struct{}"
	}

	return strings.Join([]string{"DeleteDcPointsRequest", string(data)}, " ")
}
