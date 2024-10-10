package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecoveryTimeWindowResponse Response Object
type ShowRecoveryTimeWindowResponse struct {

	// 恢复时间窗左边界（不包含）
	RecoveryMinTime *string `json:"recovery_min_time,omitempty"`

	// 恢复时间窗右边界（包含）
	RecoveryMaxTime *string `json:"recovery_max_time,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowRecoveryTimeWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecoveryTimeWindowResponse struct{}"
	}

	return strings.Join([]string{"ShowRecoveryTimeWindowResponse", string(data)}, " ")
}
