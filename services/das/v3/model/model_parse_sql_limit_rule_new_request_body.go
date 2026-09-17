package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseSqlLimitRuleNewRequestBody 解析SQL限流规则请求体
type ParseSqlLimitRuleNewRequestBody struct {

	// 原始SQL语句
	OriginalSql string `json:"original_sql"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 是否校验SQL语句
	UseTemplate bool `json:"use_template"`

	// 是否保留操作符
	KeepOperators bool `json:"keep_operators"`

	// SQL类型
	Type string `json:"type"`
}

func (o ParseSqlLimitRuleNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseSqlLimitRuleNewRequestBody struct{}"
	}

	return strings.Join([]string{"ParseSqlLimitRuleNewRequestBody", string(data)}, " ")
}
