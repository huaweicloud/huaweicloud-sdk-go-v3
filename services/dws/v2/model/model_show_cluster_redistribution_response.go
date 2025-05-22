package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterRedistributionResponse Response Object
type ShowClusterRedistributionResponse struct {
	RedisInfo *RdsRedisInfo `json:"redis_info,omitempty"`

	// **参数解释**： 调度模式。 **取值范围**： false：非调度模式；true：调度模式。
	ScheduleMode *bool `json:"schedule_mode,omitempty"`

	// **参数解释**： 是否允许暂停。 **取值范围**： false：不允许暂停；true：允许暂停。
	PauseEnable *bool `json:"pause_enable,omitempty"`

	// **参数解释**： 是否允许恢复。 **取值范围**： false：不允许恢复；true：允许恢复。
	RecoverEnable *bool `json:"recover_enable,omitempty"`

	// **参数解释**： 是否允许重试。 **取值范围**： false：不允许重试；true：允许重试。
	RetryEnable *bool `json:"retry_enable,omitempty"`

	// **参数解释**： 是否允许更新。 **取值范围**： false：不允许更新；true：允许更新，该参数仅仅适用于调度模式。
	UpdateEnable *bool `json:"update_enable,omitempty"`

	// **参数解释**： 是否允许控制。 **取值范围**： false：不允许控制；true：允许控制，该参数仅仅适用于调度模式。
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
