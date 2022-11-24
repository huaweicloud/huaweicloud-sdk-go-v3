package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 添加终端节点服务白名单请求体。
type EpsAddPermissionRequest struct {

	// 权限格式为：iam:domain::domain_id其中， ● “iam:domain::”为固定格式。 ● “domain_id”为可连接用户的帐号ID。 支持输入1~64个字符，包括“a~z”、“A~Z”、“0~9”或者“”。 “”表示所有终端节点可连接。 例如：iam:domain::6e9dfd51d1124e8d8498dce894923a0dd
	Permission string `json:"permission"`

	// 终端节点服务白名单描述
	Description string `json:"description"`
}

func (o EpsAddPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsAddPermissionRequest struct{}"
	}

	return strings.Join([]string{"EpsAddPermissionRequest", string(data)}, " ")
}
