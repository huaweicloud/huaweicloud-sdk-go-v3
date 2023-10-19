package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAccessControlRequest Request Object
type SwitchAccessControlRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	Body *SwitchAccessControlRequestBody `json:"body,omitempty"`
}

func (o SwitchAccessControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAccessControlRequest struct{}"
	}

	return strings.Join([]string{"SwitchAccessControlRequest", string(data)}, " ")
}
