package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SQL作业运行结果。
type SqlJobQueryDataset struct {

	// 作业类型。DDL, DCL, IMPORT, EXPORT, QUERY, INSERT.
	SqlType string `json:"sql_type" xml:"sql_type"`

	// 当语句类型为DDL时，返回其结果的列名称及类型。
	Schema *[]interface{} `json:"schema,omitempty" xml:"schema"`

	// 当语句类型为DDL时，直接返回其执行结果。
	Rows *[]interface{} `json:"rows,omitempty" xml:"rows"`
}

func (o SqlJobQueryDataset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobQueryDataset struct{}"
	}

	return strings.Join([]string{"SqlJobQueryDataset", string(data)}, " ")
}
