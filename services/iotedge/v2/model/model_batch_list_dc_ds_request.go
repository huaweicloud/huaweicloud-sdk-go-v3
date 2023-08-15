package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListDcDsRequest Request Object
type BatchListDcDsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 数据源所属的模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认值为10，取值区间为1-1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o BatchListDcDsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListDcDsRequest struct{}"
	}

	return strings.Join([]string{"BatchListDcDsRequest", string(data)}, " ")
}
