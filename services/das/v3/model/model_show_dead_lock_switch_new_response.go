package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockSwitchNewResponse Response Object
type ShowDeadLockSwitchNewResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 开关状态
	SwitchOn *bool `json:"switch_on,omitempty"`

	// 保存时长
	RetentionHours *int64 `json:"retention_hours,omitempty"`

	// 是否需要重试
	Retry *bool `json:"retry,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 是否可以开启
	CanOpen *bool `json:"can_open,omitempty"`

	// 无法开启原因
	CantOpenMsg    *string `json:"cant_open_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDeadLockSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"ShowDeadLockSwitchNewResponse", string(data)}, " ")
}
