package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPointTemplateRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`
}

func (o ShowPointTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPointTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowPointTemplateRequest", string(data)}, " ")
}
