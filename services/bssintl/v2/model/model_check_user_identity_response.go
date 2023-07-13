package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckUserIdentityResponse Response Object
type CheckUserIdentityResponse struct {

	// 状态码。具体请参考状态码。只有失败才会返回这个参数。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述信息。只有失败才会返回这个参数。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// available：该登录名称/手机号/邮箱有效。used_by_user：该登录名称/手机号/邮箱已被占用。
	CheckResult    *string `json:"check_result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckUserIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckUserIdentityResponse struct{}"
	}

	return strings.Join([]string{"CheckUserIdentityResponse", string(data)}, " ")
}
