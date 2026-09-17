package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KillProcessHistoryInfo KillProcessHistoryInfo对象
type KillProcessHistoryInfo struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 会话ID
	SessionId *int64 `json:"session_id,omitempty"`

	// 数据库用户
	User *string `json:"user,omitempty"`

	// 数据库主机
	Host *string `json:"host,omitempty"`

	// 数据库名称
	Db *string `json:"db,omitempty"`

	// 命令类型
	Command *string `json:"command,omitempty"`

	// 执行时间
	Time *int64 `json:"time,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 信息
	Info *string `json:"info,omitempty"`

	// Kill时间
	KillTime *int64 `json:"kill_time,omitempty"`

	// 会话被查杀的来源
	KilledSource *string `json:"killed_source,omitempty"`
}

func (o KillProcessHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KillProcessHistoryInfo struct{}"
	}

	return strings.Join([]string{"KillProcessHistoryInfo", string(data)}, " ")
}
