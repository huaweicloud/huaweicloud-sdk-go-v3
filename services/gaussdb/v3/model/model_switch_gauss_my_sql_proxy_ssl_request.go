package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxySslRequest Request Object
type SwitchGaussMySqlProxySslRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID。
	ProxyId string `json:"proxy_id"`

	Body *SwitchProxySslRequest `json:"body,omitempty"`
}

func (o SwitchGaussMySqlProxySslRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxySslRequest struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxySslRequest", string(data)}, " ")
}
