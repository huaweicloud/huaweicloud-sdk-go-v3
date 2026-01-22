package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserResponse Response Object
type UpdateUserResponse struct {

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	AccessKey *string `json:"access_key,omitempty"`

	// **参数解释**： 密钥。 **取值范围**： - 8-32个字符。 - 至少包含以下字符中的3种：    - 大写字母    - 小写字母    - 数字    - 特殊字符`~!@#$%^&*()-_=+\\\\|[{}];:\\'\\\",<.>/?。 - 不能与名称或倒序的名称相同。
	SecretKey *string `json:"secret_key,omitempty"`

	// **参数解释**： 需要配置权限的Vhost，一个用户可以配置多个Vhost下的资源权限。
	Vhosts         *[]AmqpUserPerm `json:"vhosts,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserResponse", string(data)}, " ")
}
