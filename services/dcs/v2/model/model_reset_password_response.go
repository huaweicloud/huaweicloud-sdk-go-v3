package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPasswordResponse Response Object
type ResetPasswordResponse struct {

	// 密码验证剩余次数
	RetryTimesLeft *string `json:"retry_times_left,omitempty"`

	// 锁定时间
	LockTime *string `json:"lock_time,omitempty"`

	// 锁定剩余时间
	LockTimeLeft   *string `json:"lock_time_left,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetPasswordResponse", string(data)}, " ")
}
