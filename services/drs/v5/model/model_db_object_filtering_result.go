package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbObjectFilteringResult 数据库表过滤规则响应体
type DbObjectFilteringResult struct {

	// 数据库库名。
	DbName *string `json:"db_name,omitempty"`

	// 数据库Schema名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 数据库表名称。
	TableName *string `json:"table_name,omitempty"`

	// 数据过滤校验结果。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 当数据过滤校验结果是false，返回校验失败的原因。
	Message *string `json:"message,omitempty"`
}

func (o DbObjectFilteringResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbObjectFilteringResult struct{}"
	}

	return strings.Join([]string{"DbObjectFilteringResult", string(data)}, " ")
}
