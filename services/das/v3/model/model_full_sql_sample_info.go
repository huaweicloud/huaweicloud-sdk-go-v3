package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlSampleInfo FullSqlSampleInfo对象
type FullSqlSampleInfo struct {

	// SQL
	Sql *string `json:"sql,omitempty"`

	// SQL模板
	SqlTemplateId *string `json:"sql_template_id,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 客户端地址
	Client *string `json:"client,omitempty"`

	// 用户名
	User *string `json:"user,omitempty"`

	// 执行时间
	ExecuteAt *int64 `json:"execute_at,omitempty"`

	// 执行耗时（ms）
	QueryTime *float64 `json:"query_time,omitempty"`

	// 锁等待时间（ms）
	LockTime *float64 `json:"lock_time,omitempty"`
}

func (o FullSqlSampleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlSampleInfo struct{}"
	}

	return strings.Join([]string{"FullSqlSampleInfo", string(data)}, " ")
}
