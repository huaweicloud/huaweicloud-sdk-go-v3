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
	LockTimeLeft *string `json:"lock_time_left,omitempty"`

	// **参数解释**： 重置密码结果编号。 **取值范围**： - 1：重置密码成功。 - 3：实例被锁定。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 重置密码结果信息。 **取值范围**： - success - instance is locked
	Message *string `json:"message,omitempty"`

	// **参数解释**： 重置密码错误信息，若重置密码成功，则为null。 **取值范围**： 不涉及。
	ExtMessage     *string `json:"ext_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetPasswordResponse", string(data)}, " ")
}
