package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificatePrivateKeyEchoResponse Response Object
type ShowCertificatePrivateKeyEchoResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	// 证书回显开关，项目粒度的,默认情况下,\"private_key_echo\"是true,证书的返回体中展示私钥。 当值为false时,证书的返回体中不展示私钥。
	PrivateKeyEcho *bool `json:"private_key_echo,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowCertificatePrivateKeyEchoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificatePrivateKeyEchoResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificatePrivateKeyEchoResponse", string(data)}, " ")
}
