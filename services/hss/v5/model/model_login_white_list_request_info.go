package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginWhiteListRequestInfo 登录白名单
type LoginWhiteListRequestInfo struct {

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 登录源IP **取值范围**： 字符长度1-256位
	LoginIp *string `json:"login_ip,omitempty"`

	// **参数解释**： 登录用户名 **取值范围**： 字符长度1-256位
	LoginUserName *string `json:"login_user_name,omitempty"`
}

func (o LoginWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"LoginWhiteListRequestInfo", string(data)}, " ")
}
