package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobBean struct {

	// 任务ID,异步查询标识
	JobId string `json:"job_id"`

	// 状态
	Status string `json:"status"`

	// 类型
	JobType string `json:"job_type"`

	// 虚拟机ID
	ServerId string `json:"server_id"`

	// 虚拟机名称
	ServerName string `json:"server_name"`

	// 开始时间
	BeginTime int64 `json:"begin_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`

	// 计费模式
	ChargeMode string `json:"charge_mode"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 双机实例HA共用的id
	HaId *string `json:"ha_id,omitempty"`

	// HA别名
	HaName *string `json:"ha_name,omitempty"`
}

func (o JobBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBean struct{}"
	}

	return strings.Join([]string{"JobBean", string(data)}, " ")
}
