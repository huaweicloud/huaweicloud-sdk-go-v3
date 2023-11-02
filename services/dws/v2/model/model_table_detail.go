package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableDetail 恢复表信息
type TableDetail struct {

	// schema名称
	SchemaName string `json:"schema_name"`

	// 表名称
	TableName string `json:"table_name"`
}

func (o TableDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableDetail struct{}"
	}

	return strings.Join([]string{"TableDetail", string(data)}, " ")
}
