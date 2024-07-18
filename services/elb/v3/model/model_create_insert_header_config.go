package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInsertHeaderConfig 要添加请求头参数。
type CreateInsertHeaderConfig struct {

	// 支持1-40个字符（不区分大小写）。 支持字母a-z,数字，短划线-和下划线_。 key不能是以下字符： connection、upgrade、content-length、transfer-encoding、keep-alive、te、host、cookie、remoteip、authority、x-forwarded-host、x-forwarded-for、x-forwarded-for-port、x-forwarded-tls-certificate-id、x-forwarded-tls-protocol、x-forwarded-tls-cipher、x-forwarded-elb-ip、x-forwarded-port、x-forwarded-elb-id、x-forwarded-elb-vip、x-real-ip、x-forwarded-proto、x-nuwa-trace-ne-in、x-nuwa-trace-ne-out。
	Key string `json:"key"`

	// 可选值：USER_DEFINED,REFERENCE_HEADER,SYSTEM_DEFINED。
	ValueType string `json:"value_type"`

	// header值。当value_type为SYSTEM_DEFINED时，value只可从CLIENT-PORT,CLIENT-IP, ELB-PROTOCOL, ELB-ID, ELB-PORT, ELB-EIP, ELB-VIP中取值。 取值范围：1-128个字符， 支持ascii码值32<=ch<=127范围内可打印字符，*和英文问号?。不能以空格开头或结尾。
	Value string `json:"value"`
}

func (o CreateInsertHeaderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInsertHeaderConfig struct{}"
	}

	return strings.Join([]string{"CreateInsertHeaderConfig", string(data)}, " ")
}
