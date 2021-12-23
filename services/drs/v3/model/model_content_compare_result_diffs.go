package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ContentCompareResultDiffs struct {
	// 源库名称。

	SourceDbName string `json:"source_db_name"`
	// 源库的表名称。

	SourceTableName string `json:"source_table_name"`
	// 内容对比结果差异。

	ContentCompareDiff []ContentCompareDiff `json:"ContentCompareDiff"`
	// 内容对比结果差异总数。

	ContentCompareDiffCount int32 `json:"content_compare_diff_count"`
}

func (o ContentCompareResultDiffs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResultDiffs struct{}"
	}

	return strings.Join([]string{"ContentCompareResultDiffs", string(data)}, " ")
}
