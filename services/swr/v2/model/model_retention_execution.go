package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetentionExecution struct {

	// 老化策略执行记录ID
	Id int32 `json:"id"`

	// 老化策略ID
	PolicyId int32 `json:"policy_id"`

	// 是否模拟运行
	DryRun bool `json:"dry_run"`

	// 执行状态，InProgress Succeed Failed Stopped
	Status string `json:"status"`

	// 触发类型，scheduled manual
	Trigger string `json:"trigger"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 结束时间
	EndTime string `json:"end_time"`
}

func (o RetentionExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetentionExecution struct{}"
	}

	return strings.Join([]string{"RetentionExecution", string(data)}, " ")
}
