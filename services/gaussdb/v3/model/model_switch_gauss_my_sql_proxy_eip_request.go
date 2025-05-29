package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxyEipRequest Request Object
type SwitchGaussMySqlProxyEipRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，此参数是实例的唯一标识。只能由英文字母、数字组成，后缀为in07，长度为36个字符。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，此参数是代理的唯一标识。只能由英文字母、数字组成，后缀为po01，长度为36个字符。
	ProxyId string `json:"proxy_id"`

	Body *SwitchGaussMySqlProxyEipRequestBody `json:"body,omitempty"`
}

func (o SwitchGaussMySqlProxyEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyEipRequest struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyEipRequest", string(data)}, " ")
}
