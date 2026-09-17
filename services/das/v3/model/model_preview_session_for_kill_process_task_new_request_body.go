package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewSessionForKillProcessTaskNewRequestBody 预览Kill进程任务的会话请求体
type PreviewSessionForKillProcessTaskNewRequestBody struct {

	// 数据库用户
	User *string `json:"user,omitempty"`

	// 数据库主机
	Host *string `json:"host,omitempty"`

	// 数据库名称
	Db *string `json:"db,omitempty"`

	// 命令类型
	Command *string `json:"command,omitempty"`

	// 会话执行时间
	Time int64 `json:"time"`

	// SQL信息
	Info *string `json:"info,omitempty"`

	// 任务持续时间
	TaskDuration int64 `json:"task_duration"`

	// 任务类型
	TaskType string `json:"task_type"`
}

func (o PreviewSessionForKillProcessTaskNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewSessionForKillProcessTaskNewRequestBody struct{}"
	}

	return strings.Join([]string{"PreviewSessionForKillProcessTaskNewRequestBody", string(data)}, " ")
}
