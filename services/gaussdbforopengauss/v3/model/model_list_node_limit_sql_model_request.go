package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodeLimitSqlModelRequest Request Object
type ListNodeLimitSqlModelRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 节点id。
	NodeId string `json:"node_id"`

	// sql模板。
	SqlModel *string `json:"sql_model,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  取值范围：0 - 10000
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNodeLimitSqlModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeLimitSqlModelRequest struct{}"
	}

	return strings.Join([]string{"ListNodeLimitSqlModelRequest", string(data)}, " ")
}
