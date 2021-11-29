package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SystemSecurityPolicy struct {
	// 系统安全策略的名称。

	Name string `json:"name"`
	// 系统安全策略的TLS协议列表。

	Protocols []string `json:"protocols"`
	// 系统安全策略的加密套件列表。

	Ciphers []string `json:"ciphers"`
	// 项目id。

	ProjectId string `json:"project_id"`
}

func (o SystemSecurityPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemSecurityPolicy struct{}"
	}

	return strings.Join([]string{"SystemSecurityPolicy", string(data)}, " ")
}
