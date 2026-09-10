package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAutoSqlLimitingResponse Response Object
type QueryAutoSqlLimitingResponse struct {

	// 限流策略CPU利用率。
	CpuUsage *int32 `json:"cpu_usage,omitempty"`

	// 限流策略活跃会话数。
	ActiveSessions *int32 `json:"active_sessions,omitempty"`

	// 限流策略CPU利用率和活跃会话数的关联关系。取值范围：and、or。
	Condition *string `json:"condition,omitempty"`

	// 限流策略满足限流条件的事件持续时间（分钟）。
	Duration *int32 `json:"duration,omitempty"`

	// 自治限流规则每天生效开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 自治限流规则每天生效结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 允许的会话数。
	SessionAllow *int32 `json:"session_allow,omitempty"`

	// 限流规则适用的用户列表。
	User *[]string `json:"user,omitempty"`

	// 限流规则适用的数据库列表。
	Db *[]string `json:"db,omitempty"`

	// 每次最大限流时长（分钟）。
	ClearTime *int32 `json:"clear_time,omitempty"`

	// 是否启用自治限流规则。
	Enable *bool `json:"enable,omitempty"`

	// 是否为关键字限流。
	IsKeyword *bool `json:"is_keyword,omitempty"`

	// 最大并发数。
	MaxConcurrency *int32 `json:"max_concurrency,omitempty"`

	// 是否保留SQL限流规则。
	RetainSqlRule *bool `json:"retain_sql_rule,omitempty"`

	// 是否开启kill会话开关。
	KillSessionSwitch *bool `json:"kill_session_switch,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o QueryAutoSqlLimitingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAutoSqlLimitingResponse struct{}"
	}

	return strings.Join([]string{"QueryAutoSqlLimitingResponse", string(data)}, " ")
}
