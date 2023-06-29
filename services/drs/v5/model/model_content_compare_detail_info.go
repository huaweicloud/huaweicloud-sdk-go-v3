package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContentCompareDetailInfo 内容对比详情信息体。
type ContentCompareDetailInfo struct {

	// 源库库名。
	SourceDb *string `json:"source_db,omitempty"`

	// 目标库库名。
	TargetDb *string `json:"target_db,omitempty"`

	// 源库表名。
	SourceTableName *string `json:"source_table_name,omitempty"`

	// 目标库表名。
	TargetTableName *string `json:"target_table_name,omitempty"`

	// 源库表行数。
	SourceRowNum *int64 `json:"source_row_num,omitempty"`

	// 目标库表行数。
	TargetRowNum *int64 `json:"target_row_num,omitempty"`

	// 对比不一致行数。
	DifferenceRowNum *int64 `json:"difference_row_num,omitempty"`

	// 行对比结果。取值： - true：一致。 - false：不一致。
	LineCompareResult *bool `json:"line_compare_result,omitempty"`

	// 内容对比结果。取值： - true：一致。 - false：不一致。
	ContentCompareResult *bool `json:"content_compare_result,omitempty"`

	// 失败原因。
	Message *string `json:"message,omitempty"`
}

func (o ContentCompareDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareDetailInfo struct{}"
	}

	return strings.Join([]string{"ContentCompareDetailInfo", string(data)}, " ")
}
