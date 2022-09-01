package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAccessPolicyRespPolicies struct {

	// 秘钥。
	AccessKey *string `json:"access_key,omitempty" xml:"access_key"`

	// IP白名单。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty" xml:"white_remote_address"`

	// 是否为管理员。
	Admin *bool `json:"admin,omitempty" xml:"admin"`

	// 权限。
	Perm *string `json:"perm,omitempty" xml:"perm"`
}

func (o ListAccessPolicyRespPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyRespPolicies struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyRespPolicies", string(data)}, " ")
}
