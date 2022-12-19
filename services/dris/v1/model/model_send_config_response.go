package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：分发通道选择及目标推送设备的配置。
type SendConfigResponse struct {
	Channel *Channel `json:"channel,omitempty"`

	TargetList *TargetList `json:"target_list,omitempty"`
}

func (o SendConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendConfigResponse struct{}"
	}

	return strings.Join([]string{"SendConfigResponse", string(data)}, " ")
}
