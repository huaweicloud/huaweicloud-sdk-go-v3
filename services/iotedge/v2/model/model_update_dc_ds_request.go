package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDcDsRequest Request Object
type UpdateDcDsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	// 指此配置是否只更新了名称，默认值为false。 - true: 配置中只更新了名称 - false: 配置中包含其他配置参数更新
	UpdateNameOnly *bool `json:"update_name_only,omitempty"`

	Body *UpdateDcDsReqDto `json:"body,omitempty"`
}

func (o UpdateDcDsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcDsRequest struct{}"
	}

	return strings.Join([]string{"UpdateDcDsRequest", string(data)}, " ")
}
