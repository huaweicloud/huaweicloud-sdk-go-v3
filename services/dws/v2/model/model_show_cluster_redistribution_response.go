package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterRedistributionResponse Response Object
type ShowClusterRedistributionResponse struct {
	RedisInfo *RdsRedisInfo `json:"redis_info,omitempty"`

	// 调度模式
	ScheduleMode *bool `json:"schedule_mode,omitempty"`

	// 是否允许暂停
	PauseEnable *bool `json:"pause_enable,omitempty"`

	// 是否允许恢复
	RecoverEnable *bool `json:"recover_enable,omitempty"`

	// 是否允许重试
	RetryEnable *bool `json:"retry_enable,omitempty"`

	// 是否允许更新
	UpdateEnable *bool `json:"update_enable,omitempty"`

	// 是否允许控制
	ControlEnable  *bool `json:"control_enable,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowClusterRedistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRedistributionResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterRedistributionResponse", string(data)}, " ")
}
