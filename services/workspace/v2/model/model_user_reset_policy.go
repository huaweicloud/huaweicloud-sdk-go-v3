package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserResetPolicy 用户重置策略。
type UserResetPolicy struct {

	// 开关
	Enable *bool `json:"enable,omitempty"`

	// 用户每天自动重试次数,重置次数每天0点刷新。
	ResetCountPerDay *int32 `json:"reset_count_per_day,omitempty"`
}

func (o UserResetPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserResetPolicy struct{}"
	}

	return strings.Join([]string{"UserResetPolicy", string(data)}, " ")
}
