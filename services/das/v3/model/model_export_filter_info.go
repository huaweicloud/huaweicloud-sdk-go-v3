package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFilterInfo 导出筛选条件
type ExportFilterInfo struct {

	// 数据库名称列表
	DbNames *[]string `json:"db_names,omitempty"`

	// 表名称列表
	TbNames *[]string `json:"tb_names,omitempty"`

	// 文件名称列表
	FileNames *[]string `json:"file_names,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// SQL类型列表（insert、update、delete、ddl）
	TypeList *[]string `json:"type_list,omitempty"`

	// 列列表
	ColumnList *[]ExportColumnInfo `json:"column_list,omitempty"`

	// 是否将UPDATE语句导出为两条INSERT语句
	ParseDoubleInsert *bool `json:"parse_double_insert,omitempty"`
}

func (o ExportFilterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFilterInfo struct{}"
	}

	return strings.Join([]string{"ExportFilterInfo", string(data)}, " ")
}
