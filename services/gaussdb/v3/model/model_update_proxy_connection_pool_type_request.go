package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProxyConnectionPoolTypeRequest struct {

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ProxyUpdateProxyConnectionPoolTypeRequest `json:"body,omitempty"`
}

func (o UpdateProxyConnectionPoolTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyConnectionPoolTypeRequest struct{}"
	}

	return strings.Join([]string{"UpdateProxyConnectionPoolTypeRequest", string(data)}, " ")
}
