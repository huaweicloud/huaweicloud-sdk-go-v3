package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapTableColumnTreeNodeTableColumnDto 表格列树节点数据传输对象
type IsapTableColumnTreeNodeTableColumnDto struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 父级名称
	ParentName *string `json:"parent_name,omitempty"`

	// 所属名称
	OwnName *string `json:"own_name,omitempty"`

	// 深度
	Depth *int32 `json:"depth,omitempty"`

	Source *IsapTableColumnDto `json:"source,omitempty"`

	// 子节点数组
	Children *[]IsapTableColumnTreeNodeTableColumnDto `json:"children,omitempty"`

	// 父级名称
	PreviousName *string `json:"previous_name,omitempty"`

	// 列序号
	ColumnSequenceNumber *int32 `json:"column_sequence_number,omitempty"`
}

func (o IsapTableColumnTreeNodeTableColumnDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableColumnTreeNodeTableColumnDto struct{}"
	}

	return strings.Join([]string{"IsapTableColumnTreeNodeTableColumnDto", string(data)}, " ")
}
