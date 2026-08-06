package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchFullsqlSwitchResponse Response Object
type SwitchFullsqlSwitchResponse struct {

	// 开关状态
	OpenStatus *int32 `json:"open_status,omitempty"`

	// 保留天数
	RetentionDays *int32 `json:"retention_days,omitempty"`

	// 是否能开启
	CanOpen *bool `json:"can_open,omitempty"`

	// 不能开启的原因
	CantOpenMsg *string `json:"cant_open_msg,omitempty"`

	// 上次开启时间
	LastOpenTime   float32 `json:"last_open_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchFullsqlSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFullsqlSwitchResponse struct{}"
	}

	return strings.Join([]string{"SwitchFullsqlSwitchResponse", string(data)}, " ")
}
