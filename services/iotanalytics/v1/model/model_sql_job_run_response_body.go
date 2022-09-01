package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SQL作业运行响应。
type SqlJobRunResponseBody struct {

	// 作业类型。DDL, DCL, IMPORT, EXPORT, QUERY, INSERT.
	SqlType string `json:"sql_type" xml:"sql_type"`

	// 当语句类型为DDL时，返回其结果的列名称及类型。
	Schema *[]interface{} `json:"schema,omitempty" xml:"schema"`

	// 当语句类型为DDL时，直接返回其执行结果。
	Rows *[]interface{} `json:"rows,omitempty" xml:"rows"`

	// 作业执行模式：async: 异步; sync: 同步。
	JobMode *string `json:"job_mode,omitempty" xml:"job_mode"`
}

func (o SqlJobRunResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobRunResponseBody struct{}"
	}

	return strings.Join([]string{"SqlJobRunResponseBody", string(data)}, " ")
}
