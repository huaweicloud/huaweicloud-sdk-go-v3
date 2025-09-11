package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Job struct {

	// 任务ID
	Id int32 `json:"id"`

	// 触发器ID
	PolicyId int32 `json:"policy_id"`

	// 事件类型
	EventType string `json:"event_type"`

	// 通知类型
	NotifyType string `json:"notify_type"`

	// 任务状态
	Status string `json:"status"`

	// 状态原因
	StatusText string `json:"status_text"`

	// 任务详情
	JobDetail string `json:"job_detail"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
