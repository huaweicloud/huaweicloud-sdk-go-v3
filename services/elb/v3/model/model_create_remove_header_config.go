package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRemoveHeaderConfig struct {

	// **参数解释**：被移除的请求头的参数名。  **约束限制**：不能移除以下请求头参数： connection、upgrade、content-length、transfer-encoding、keep-alive、te、host、cookie、remoteip、authority、x-forwarded-host、x-forwarded-for、x-forwarded-for-port、x-forwarded-tls-certificate-id、x-forwarded-tls-protocol、x-forwarded-tls-cipher、x-forwarded-elb-ip、x-forwarded-port、x-forwarded-elb-id、x-forwarded-elb-vip、x-real-ip、x-forwarded-proto、x-nuwa-trace-ne-in、x-nuwa-trace-ne-out  **取值范围**：1-40个字符，字母a-z（不区分大小写）、数字，短划线-和下划线_。  **默认取值**：不涉及
	Key string `json:"key"`
}

func (o CreateRemoveHeaderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemoveHeaderConfig struct{}"
	}

	return strings.Join([]string{"CreateRemoveHeaderConfig", string(data)}, " ")
}
