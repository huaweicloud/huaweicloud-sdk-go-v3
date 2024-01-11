package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanLog 工作计划日志
type PlanLog struct {

	// 执行时间
	ExecTime string `json:"exec_time"`

	// 执行计划阶段
	StageInfo string `json:"stage_info"`

	// 执行结果。
	ExecResult int32 `json:"exec_result"`

	// 执行日志。
	ExecLog string `json:"exec_log"`
}

func (o PlanLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanLog struct{}"
	}

	return strings.Join([]string{"PlanLog", string(data)}, " ")
}
