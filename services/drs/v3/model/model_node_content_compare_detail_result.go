package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeContentCompareDetailResult 内容对比详情
type NodeContentCompareDetailResult struct {

	// 源库名称。
	SourceDb *string `json:"source_db,omitempty"`

	// 目标库名称。
	TargetDb *string `json:"target_db,omitempty"`

	// 源库的表名称。
	SourceTableName *string `json:"source_table_name,omitempty"`

	// 目标库名称。
	TargetTableName *string `json:"target_table_name,omitempty"`

	// 源库表行数。
	SourceRowNum *int64 `json:"source_row_num,omitempty"`

	// 目标库表行数。
	TargetRowNum *int64 `json:"target_row_num,omitempty"`

	// 源库的表和目标库的表的差异值。
	DifferenceRowNum *int64 `json:"difference_row_num,omitempty"`

	// 行对比结果。 - true：一致 - false：不一致
	LineCompareResult *bool `json:"line_compare_result,omitempty"`

	// 内容对比结果。 - true：一致 - false：不一致
	ContentCompareResult *bool `json:"content_compare_result,omitempty"`

	// 附加信息。
	Message *string `json:"message,omitempty"`

	// 行过滤配置条件
	CompareLineConfigFilter *string `json:"compare_line_config_filter,omitempty"`
}

func (o NodeContentCompareDetailResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeContentCompareDetailResult struct{}"
	}

	return strings.Join([]string{"NodeContentCompareDetailResult", string(data)}, " ")
}
