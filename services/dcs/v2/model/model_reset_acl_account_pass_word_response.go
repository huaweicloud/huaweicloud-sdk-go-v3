package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetAclAccountPassWordResponse Response Object
type ResetAclAccountPassWordResponse struct {

	// 任务执行成功代码。
	Code *string `json:"code,omitempty"`

	// 重置结果说明信息。
	Message *string `json:"message,omitempty"`

	// 锁定时间。验证失败时和锁定时该参数返回不为null
	LockTime *string `json:"lock_time,omitempty"`

	// 锁定剩余时间。锁定时该参数返回不为null
	LockTimeLeft *string `json:"lock_time_left,omitempty"`

	// 密码验证剩余次数。验证失败时该参数返回不为null
	RetryTimesLeft *string `json:"retry_times_left,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetAclAccountPassWordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAclAccountPassWordResponse struct{}"
	}

	return strings.Join([]string{"ResetAclAccountPassWordResponse", string(data)}, " ")
}
