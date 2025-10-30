package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecommendSqlLimitRuleRespRawSql mysql 提供原始sql的信息， taurus 只提供一条sql信息
type RecommendSqlLimitRuleRespRawSql struct {

	// 会话id
	SessionId *string `json:"session_id,omitempty"`

	// 主机ip
	Host *string `json:"host,omitempty"`

	// sql
	Sql *string `json:"sql,omitempty"`

	// 数据库名称
	Db *string `json:"db,omitempty"`

	// 时间
	Time *int64 `json:"time,omitempty"`
}

func (o RecommendSqlLimitRuleRespRawSql) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecommendSqlLimitRuleRespRawSql struct{}"
	}

	return strings.Join([]string{"RecommendSqlLimitRuleRespRawSql", string(data)}, " ")
}
