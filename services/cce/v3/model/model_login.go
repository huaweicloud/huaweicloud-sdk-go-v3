package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Login struct {

	// 选择密钥对方式登录时的密钥对名称。
	SshKey *string `json:"sshKey,omitempty"`

	UserPassword *UserPassword `json:"userPassword,omitempty"`

	// **参数解释**： 更新节点池时，移除当前节点池密码方式登录的配置 **约束限制**： 仅更新节点池场景支持该参数，设置为true时不允许设置userPassword **取值范围**： 不涉及 **默认取值**： false
	RemoveUserPassword *bool `json:"removeUserPassword,omitempty"`

	// **参数解释**： 更新节点池时，移除当前节点池密钥对方式登录的配置 **约束限制**： 仅更新节点池场景支持该参数，设置为true时不允许设置sshKey **取值范围**： 不涉及 **默认取值**： false
	RemoveSSHKey *bool `json:"removeSSHKey,omitempty"`
}

func (o Login) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Login struct{}"
	}

	return strings.Join([]string{"Login", string(data)}, " ")
}
