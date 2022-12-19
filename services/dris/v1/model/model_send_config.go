package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：分发通道及目标设备。
type SendConfig struct {
	Channel *Channel `json:"channel"`

	TargetList *TargetList `json:"target_list,omitempty"`
}

func (o SendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendConfig struct{}"
	}

	return strings.Join([]string{"SendConfig", string(data)}, " ")
}
