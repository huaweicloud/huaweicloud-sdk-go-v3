package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInsertHeaderConfig 要添加请求头参数。
type UpdateInsertHeaderConfig struct {

	// 参数解释：请求头参数名。  约束限制：不能是以下字符： connection、upgrade、content-length、transfer-encoding、keep-alive、te、host、cookie、remoteip、authority、x-forwarded-host、x-forwarded-for、x-forwarded-for-port、x-forwarded-tls-certificate-id、x-forwarded-tls-protocol、x-forwarded-tls-cipher、x-forwarded-elb-ip、x-forwarded-port、x-forwarded-elb-id、x-forwarded-elb-vip、x-real-ip、x-forwarded-proto、x-nuwa-trace-ne-in、x-nuwa-trace-ne-out。  取值范围：1-40个字符，字母a-z（不区分大小写）、数字，短划线-和下划线_。
	Key string `json:"key"`

	// 参数解释：请求头参数类别。  取值范围：USER_DEFINED,REFERENCE_HEADER,SYSTEM_DEFINED。
	ValueType string `json:"value_type"`

	// 参数解释：请求头参数的值。  约束限制：当value_type为SYSTEM_DEFINED时，value只可从CLIENT-PORT,CLIENT-IP, ELB-PROTOCOL, ELB-ID, ELB-PORT, ELB-EIP, ELB-VIP中取值。  取值范围：1-128个字符，支持ascii码值32<=ch<=127范围内可打印字符，*和英文问号?。不能以空格开头或结尾。
	Value string `json:"value"`
}

func (o UpdateInsertHeaderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInsertHeaderConfig struct{}"
	}

	return strings.Join([]string{"UpdateInsertHeaderConfig", string(data)}, " ")
}
