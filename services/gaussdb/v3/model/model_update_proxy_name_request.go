package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyNameRequest Request Object
type UpdateProxyNameRequest struct {

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	Body *ProxyUpdateProxyNameRequest `json:"body,omitempty"`
}

func (o UpdateProxyNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateProxyNameRequest", string(data)}, " ")
}
