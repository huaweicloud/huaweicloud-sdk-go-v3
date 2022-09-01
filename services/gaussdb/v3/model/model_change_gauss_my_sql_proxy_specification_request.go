package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeGaussMySqlProxySpecificationRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id" xml:"proxy_id"`

	Body *TaurusProxyScaleRequest `json:"body,omitempty" xml:"body"`
}

func (o ChangeGaussMySqlProxySpecificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeGaussMySqlProxySpecificationRequest struct{}"
	}

	return strings.Join([]string{"ChangeGaussMySqlProxySpecificationRequest", string(data)}, " ")
}
