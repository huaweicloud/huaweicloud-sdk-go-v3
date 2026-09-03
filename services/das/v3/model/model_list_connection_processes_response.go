package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionProcessesResponse Response Object
type ListConnectionProcessesResponse struct {

	// 会话信息列表
	ProcessInfoList *[]ProcessInfo `json:"process_info_list,omitempty"`

	// 根据条件筛选的总会话数
	Total *int64 `json:"total,omitempty"`

	// 用户列表
	UserInfoList *[]string `json:"user_info_list,omitempty"`

	// 数据库列表
	DbInfoList *[]string `json:"db_info_list,omitempty"`

	// 来源IP列表
	HostInfoList *[]string `json:"host_info_list,omitempty"`

	// 状态列表
	StateInfoList *[]string `json:"state_info_list,omitempty"`

	// 命令列表
	CommandInfoList *[]string `json:"command_info_list,omitempty"`

	// 会话执行时间比例
	SessionExecTime *interface{} `json:"session_exec_time,omitempty"`

	// 空闲会话数
	IdleSession *int64 `json:"idle_session,omitempty"`

	// 运行会话数
	ActiveSession *int64 `json:"active_session,omitempty"`

	// 概要
	Summary *[]ProcessSummary `json:"summary,omitempty"`

	// 按用户统计信息
	UserStats *[]ProcessStats `json:"user_stats,omitempty"`

	// 按访问来源统计
	HostStats *[]ProcessStats `json:"host_stats,omitempty"`

	// 按数据库统计
	DbStats *[]ProcessStats `json:"db_stats,omitempty"`

	// 是否显示版本支持信息
	ShowVersionSupportMessage *bool `json:"show_version_support_message,omitempty"`

	// 是否告警信息
	ShowWarnMessage *bool `json:"show_warn_message,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListConnectionProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionProcessesResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionProcessesResponse", string(data)}, " ")
}
