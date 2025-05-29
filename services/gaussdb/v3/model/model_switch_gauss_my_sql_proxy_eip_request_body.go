package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxyEipRequestBody Proxy绑定解绑弹性公网IP请求体。
type SwitchGaussMySqlProxyEipRequestBody struct {

	// 待绑定的弹性公网IP地址。
	PublicIp string `json:"public_ip"`

	// 弹性公网IP地址对应的ID。请求为绑定弹性公网IP时需传入该参数，请求为解绑弹性公网IP时无需传入。
	PublicIpId *string `json:"public_ip_id,omitempty"`

	// 请求是否为绑定弹性公网IP。  取值范围： - true：表示请求为绑定弹性公网IP。 - false：表示请求为解绑弹性公网IP。
	Bind string `json:"bind"`
}

func (o SwitchGaussMySqlProxyEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyEipRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyEipRequestBody", string(data)}, " ")
}
