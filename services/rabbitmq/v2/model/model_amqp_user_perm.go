package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AmqpUserPerm struct {

	// 需要配置权限的Vhost名称，一个用户可以配置多个Vhost下的资源权限。
	Vhost *string `json:"vhost,omitempty"`

	// 使用正则表达式匹配资源配置权限。例如，在输入框内输入“^janeway-.*”，则表示授权给该用户当前Vhost下，所有名称以“janeway-”开头的资源的配置权限。
	Conf *string `json:"conf,omitempty"`

	// 使用正则表达式匹配资源写权限。例如，在输入框内输入“.*”，则表示授权给该用户当前Vhost下，所有资源的写权限。
	Write *string `json:"write,omitempty"`

	// 使用正则表达式匹配资源读权限。例如，在输入框内输入“.*”，则表示授权给该用户当前Vhost下，所有资源的读权限。
	Read *string `json:"read,omitempty"`
}

func (o AmqpUserPerm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AmqpUserPerm struct{}"
	}

	return strings.Join([]string{"AmqpUserPerm", string(data)}, " ")
}
