package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIaConfigsRequest Request Object
type ListIaConfigsRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 边侧第三方应用的模块ID
	IaId string `json:"ia_id"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，取值范围为非负整数，默认值为10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIaConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIaConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListIaConfigsRequest", string(data)}, " ")
}
