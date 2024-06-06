package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartProxyInstanceRequest Request Object
type RestartProxyInstanceRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID。
	ProxyId string `json:"proxy_id"`
}

func (o RestartProxyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartProxyInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartProxyInstanceRequest", string(data)}, " ")
}
