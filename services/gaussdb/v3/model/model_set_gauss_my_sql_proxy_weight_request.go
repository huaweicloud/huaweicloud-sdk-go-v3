package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetGaussMySqlProxyWeightRequest Request Object
type SetGaussMySqlProxyWeightRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	Body *TaurusModifyProxyWeightRequest `json:"body,omitempty"`
}

func (o SetGaussMySqlProxyWeightRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetGaussMySqlProxyWeightRequest struct{}"
	}

	return strings.Join([]string{"SetGaussMySqlProxyWeightRequest", string(data)}, " ")
}
