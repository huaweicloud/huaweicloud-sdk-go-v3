package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDcDsRequest Request Object
type UpdateDcDsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	Body *UpdateDcDsReqDto `json:"body,omitempty"`
}

func (o UpdateDcDsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcDsRequest struct{}"
	}

	return strings.Join([]string{"UpdateDcDsRequest", string(data)}, " ")
}
