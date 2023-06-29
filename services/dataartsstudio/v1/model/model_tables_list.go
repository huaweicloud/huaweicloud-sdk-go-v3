package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TablesList schema信息
type TablesList struct {

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 表id
	TableId *string `json:"table_id,omitempty"`

	// 表的中文名称
	TableNameCn *string `json:"table_name_cn,omitempty"`

	// 表中字段
	Columns *string `json:"columns,omitempty"`

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 数据连接名称
	DwName *string `json:"dw_name,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 表的生命周期
	LifeCycle *int32 `json:"life_cycle,omitempty"`

	// 表的描述
	Description *string `json:"description,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 数据连接id
	ProjectId *string `json:"project_id,omitempty"`

	// 表的创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 表的大小
	TableSize *int32 `json:"table_size,omitempty"`

	// 当前查询条件下表的总记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 表是否合规
	IsValid *int32 `json:"is_valid,omitempty"`

	// 表的额外设置
	ExtraSetting *string `json:"extra_setting,omitempty"`

	// 是否进行数据分区
	Partitioned *bool `json:"partitioned,omitempty"`
}

func (o TablesList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TablesList struct{}"
	}

	return strings.Join([]string{"TablesList", string(data)}, " ")
}
