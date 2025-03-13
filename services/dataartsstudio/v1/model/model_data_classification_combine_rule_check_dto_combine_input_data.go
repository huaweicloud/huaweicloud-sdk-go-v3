package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataClassificationCombineRuleCheckDtoCombineInputData 模拟数据
type DataClassificationCombineRuleCheckDtoCombineInputData struct {

	// 字段内容模拟数据
	ColumnContent *string `json:"column_content,omitempty"`

	// 列名模拟数据
	ColumnName *string `json:"column_name,omitempty"`

	// 列注释模拟数据
	ColumnComment *string `json:"column_comment,omitempty"`

	// 表注释模拟数据
	TableComment *string `json:"table_comment,omitempty"`

	// 表名模拟数据
	TableName *string `json:"table_name,omitempty"`

	// 库名模拟数据
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o DataClassificationCombineRuleCheckDtoCombineInputData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationCombineRuleCheckDtoCombineInputData struct{}"
	}

	return strings.Join([]string{"DataClassificationCombineRuleCheckDtoCombineInputData", string(data)}, " ")
}
