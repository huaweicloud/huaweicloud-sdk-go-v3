package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlJobRun SQL作业运行。
type SqlJobRun struct {

	// 作业类型。DDL, DCL, IMPORT, EXPORT, QUERY, INSERT.
	SqlType string `json:"sql_type"`
}

func (o SqlJobRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobRun struct{}"
	}

	return strings.Join([]string{"SqlJobRun", string(data)}, " ")
}
