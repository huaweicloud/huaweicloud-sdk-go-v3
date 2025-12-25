package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableSchemaDto 表结构数据传输对象
type TableSchemaDto struct {

	// 表格列列表
	Columns *[]IsapTableColumnDto `json:"columns,omitempty"`

	// 表主键列表
	PrimaryKeyList *[]string `json:"primary_key_list,omitempty"`

	// 表分区列表
	Partition *[]string `json:"partition,omitempty"`

	// 表水位列
	WatermarkColumn *string `json:"watermark_column,omitempty"`

	// 表水位延迟间隔
	WatermarkInterval *float32 `json:"watermark_interval,omitempty"`

	// 表时间过滤列
	TimeFilter *string `json:"time_filter,omitempty"`

	// 展示设置列表
	DisplaySetting *[]IsapTableColumnDisplaySettingDto `json:"display_setting,omitempty"`

	ColumnTreeRoot *IsapTableColumnTreeNodeTableColumnDto `json:"column_tree_root,omitempty"`

	// 表主键列表
	PrimaryKey *[]string `json:"primary_key,omitempty"`

	// 分区键
	PartitionKey *string `json:"partition_key,omitempty"`
}

func (o TableSchemaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableSchemaDto struct{}"
	}

	return strings.Join([]string{"TableSchemaDto", string(data)}, " ")
}
