package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeLineageGuids 导入血缘响应体。
type NodeLineageGuids struct {

	// 节点guid。
	Node *string `json:"node,omitempty"`

	// schema名称。
	ColumnLineages *[]string `json:"column_lineages,omitempty"`
}

func (o NodeLineageGuids) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeLineageGuids struct{}"
	}

	return strings.Join([]string{"NodeLineageGuids", string(data)}, " ")
}
