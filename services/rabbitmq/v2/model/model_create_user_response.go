package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserResponse Response Object
type CreateUserResponse struct {

	// 用户名，只能英文字母开头，且由英文字母、数字、中划线、下划线组成，长度为7~64个字符。
	AccessKey *string `json:"access_key,omitempty"`

	// 密钥。 8-32个字符。 至少包含以下字符中的3种：   - 大写字母   - 小写字母   - 数字   - 特殊字符`~!@#$%^&*()-_=+\\\\|[{}];:\\'\\\",<.>/?。 不能与名称或倒序的名称相同。
	SecretKey *string `json:"secret_key,omitempty"`

	// 需要配置权限的 Vhost，一个用户可以配置多个Vhost下的资源权限。
	Vhosts         *[]AmqpUserPerm `json:"vhosts,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserResponse struct{}"
	}

	return strings.Join([]string{"CreateUserResponse", string(data)}, " ")
}
