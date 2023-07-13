package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendedKeyUsage struct {

	// 服务器身份验证，OID为：1.3.6.1.5.5.7.3.1。 - **true** - **false** > 服务器证书请启用本增强型密钥用法，默认为false。
	ServerAuth *bool `json:"server_auth,omitempty"`

	// 客户端身份验证，OID为：1.3.6.1.5.5.7.3.2。 - **true** - **false** > 客户端证书请启用本增强型密钥用法，默认为false。
	ClientAuth *bool `json:"client_auth,omitempty"`

	// 代码签名，OID为：1.3.6.1.5.5.7.3.3。 - **true** - **false** > 签署可下载的可执行代码客户端认证，默认为false。
	CodeSigning *bool `json:"code_signing,omitempty"`

	// 安全电子邮件，OID为：1.3.6.1.5.5.7.3.4。 - **true** - **false** > 电子邮件保护，默认为false。
	EmailProtection *bool `json:"email_protection,omitempty"`

	// 时间戳，OID为：1.3.6.1.5.5.7.3.8。 - **true** - **false** > 将一个对象的哈希绑定到一个时间，默认为false。
	TimeStamping *bool `json:"time_stamping,omitempty"`
}

func (o ExtendedKeyUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendedKeyUsage struct{}"
	}

	return strings.Join([]string{"ExtendedKeyUsage", string(data)}, " ")
}
