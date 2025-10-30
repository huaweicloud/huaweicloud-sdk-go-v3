package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckUserIdentityResponse Response Object
type CheckUserIdentityResponse struct {

	// 状态码。具体请参考状态码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 是否可以继续注册。该参数非必填，且只允许字符串,available: 该登录名称/手机号/邮箱可以继续注册,used_by_user: 该登录名称/手机号/邮箱不可以继续注册
	CheckResult *string `json:"check_result,omitempty"`

	// 手机号剩余可注册客户数量。该参数非必填，只有search_type=mobile时，该字段才返回值。表示可以继续使用该手机号注册客户的数量。
	MobileRemainAvailableNum *int32 `json:"mobile_remain_available_num,omitempty"`
	HttpStatusCode           int    `json:"-"`
}

func (o CheckUserIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckUserIdentityResponse struct{}"
	}

	return strings.Join([]string{"CheckUserIdentityResponse", string(data)}, " ")
}
