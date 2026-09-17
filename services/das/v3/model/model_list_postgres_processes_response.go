package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPostgresProcessesResponse Response Object
type ListPostgresProcessesResponse struct {

	// 进程列表
	ProcessInfoList *[]PgProcessInfo `json:"process_info_list,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 用户信息列表
	UserInfoList *[]string `json:"user_info_list,omitempty"`

	// 数据库信息列表
	DbInfoList *[]string `json:"db_info_list,omitempty"`

	// 主机信息列表
	HostInfoList *[]string `json:"host_info_list,omitempty"`

	// 状态信息列表
	StateInfoList *[]string `json:"state_info_list,omitempty"`

	// 命令信息列表
	CommandInfoList *[]string `json:"command_info_list,omitempty"`

	// 会话执行时间比例
	SessionExecTime *interface{} `json:"session_exec_time,omitempty"`

	// 空闲会话数
	IdleSession *int64 `json:"idle_session,omitempty"`

	// 运行会话数
	ActiveSession *int64 `json:"active_session,omitempty"`

	// 概要
	Summary *[]PgProcessSummary `json:"summary,omitempty"`

	// 按用户统计信息
	UserStats *[]PgProcessStats `json:"user_stats,omitempty"`

	// 按访问来源统计
	HostStats *[]PgProcessStats `json:"host_stats,omitempty"`

	// 按数据库统计
	DbStats *[]PgProcessStats `json:"db_stats,omitempty"`

	// 是否显示版本支持信息
	ShowVersionSupportMessage *bool `json:"show_version_support_message,omitempty"`

	// 是否显示告警信息
	ShowWarnMessage *bool `json:"show_warn_message,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListPostgresProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresProcessesResponse struct{}"
	}

	return strings.Join([]string{"ListPostgresProcessesResponse", string(data)}, " ")
}
