package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAccessPolicyRespPolicies struct {

	// 用户名。
	AccessKey *string `json:"access_key,omitempty"`

	// 密钥。
	SecretKey *string `json:"secret_key,omitempty"`

	// IP白名单。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// 是否为管理员。
	Admin *bool `json:"admin,omitempty"`

	// 权限。
	Perm *string `json:"perm,omitempty"`
}

func (o ListAccessPolicyRespPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyRespPolicies struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyRespPolicies", string(data)}, " ")
}
