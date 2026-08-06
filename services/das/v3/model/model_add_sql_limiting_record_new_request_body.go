package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSqlLimitingRecordNewRequestBody 新增SQL限流规则请求体
type AddSqlLimitingRecordNewRequestBody struct {

	// 引擎类型
	EngineType string `json:"engine_type"`

	// SQL类型
	Type string `json:"type"`

	// 限流规则
	KeyStr string `json:"key_str"`

	// 最大等待时间
	MaxWaiting *int32 `json:"max_waiting,omitempty"`

	// 最大并发数
	MaxConnection int32 `json:"max_connection"`

	// 历史会话限流开关
	HisSqlLimitingSwitch *bool `json:"his_sql_limiting_switch,omitempty"`

	// 节点ID列表
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 是否自动化
	Automaticity *bool `json:"automaticity,omitempty"`

	// 过期时间
	Duration *int32 `json:"duration,omitempty"`
}

func (o AddSqlLimitingRecordNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSqlLimitingRecordNewRequestBody struct{}"
	}

	return strings.Join([]string{"AddSqlLimitingRecordNewRequestBody", string(data)}, " ")
}
