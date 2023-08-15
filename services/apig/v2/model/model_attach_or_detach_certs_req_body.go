package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachOrDetachCertsReqBody 域名绑定和解绑证书的请求体
type AttachOrDetachCertsReqBody struct {

	// 证书的id集合
	CertificateIds []string `json:"certificate_ids"`

	// 是否开启客户端证书校验。当绑定证书存在trusted_root_ca时，默认开启；当绑定证书不存在trusted_root_ca时，默认关闭。
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`
}

func (o AttachOrDetachCertsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachOrDetachCertsReqBody struct{}"
	}

	return strings.Join([]string{"AttachOrDetachCertsReqBody", string(data)}, " ")
}
