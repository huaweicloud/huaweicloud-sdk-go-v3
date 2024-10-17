package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobBean struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 任务状态 - SUCCESS - RUNNING - FAIL - INIT - READY
	Status string `json:"status"`

	// 类型
	JobType string `json:"job_type"`

	// 虚拟机ID
	ServerId string `json:"server_id"`

	// 虚拟机名称
	ServerName string `json:"server_name"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 开始时间
	BeginTime int64 `json:"begin_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`

	// 计费模式 - Period:包周期计费 - Demand:按需计费
	ChargeMode string `json:"charge_mode"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 防护实例ID,该字段已废弃
	HaId *string `json:"ha_id,omitempty"`

	// 防护实例名称，该字段已废弃
	HaName *string `json:"ha_name,omitempty"`
}

func (o JobBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBean struct{}"
	}

	return strings.Join([]string{"JobBean", string(data)}, " ")
}
