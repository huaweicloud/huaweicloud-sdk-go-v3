package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BigKeyHotKeySwitchStatus struct {

	// 开关状态。0-关闭，1-开启
	OpenStatus *int32 `json:"open_status,omitempty"`

	// 数据保留小时数
	RetentionHours *int64 `json:"retention_hours,omitempty"`

	// 是否可以开启
	CanOpen *bool `json:"can_open,omitempty"`

	// 无法开启的原因（如内核版本不满足等）
	CantOpenMsg *string `json:"cant_open_msg,omitempty"`

	// 最后开启时间（毫秒时间戳）
	LastOpenTime *int64 `json:"last_open_time,omitempty"`
}

func (o BigKeyHotKeySwitchStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BigKeyHotKeySwitchStatus struct{}"
	}

	return strings.Join([]string{"BigKeyHotKeySwitchStatus", string(data)}, " ")
}
