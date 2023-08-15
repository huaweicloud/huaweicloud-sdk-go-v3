package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DirtyData 异常数据。
type DirtyData struct {

	// 数据库名称。
	DbName *string `json:"db_name,omitempty"`

	// 模式名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名称。
	TableName *string `json:"table_name,omitempty"`

	// 异常SQL。
	ErrorSql *string `json:"error_sql,omitempty"`

	// 发生异常时间，UTC时间，例如：2023-06-10T03:01:52Z
	ErrorTime *string `json:"error_time,omitempty"`

	// 异常信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DirtyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirtyData struct{}"
	}

	return strings.Join([]string{"DirtyData", string(data)}, " ")
}
