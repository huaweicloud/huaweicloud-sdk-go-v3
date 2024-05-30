package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyVersionRequest Request Object
type ShowProxyVersionRequest struct {

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// engine
	EngineName string `json:"engine_name"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowProxyVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowProxyVersionRequest", string(data)}, " ")
}
