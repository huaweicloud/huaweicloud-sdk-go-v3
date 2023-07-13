package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeGaussMySqlProxySpecificationRequest Request Object
type ChangeGaussMySqlProxySpecificationRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	Body *TaurusProxyScaleRequest `json:"body,omitempty"`
}

func (o ChangeGaussMySqlProxySpecificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeGaussMySqlProxySpecificationRequest struct{}"
	}

	return strings.Join([]string{"ChangeGaussMySqlProxySpecificationRequest", string(data)}, " ")
}
