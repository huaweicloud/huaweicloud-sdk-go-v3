package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设置证书请求体
type HttpPutBody struct {

	// HTTPS证书是否启用。（on：开启，off：关闭）
	HttpsStatus *string `json:"https_status,omitempty"`

	// 证书名字。（长度限制为3-32字符）。当证书开启时必传。
	CertificateName *string `json:"certificate_name,omitempty"`

	// HTTPS协议使用的证书内容，当证书开启时必传。取值范围：PEM编码格式。
	CertificateValue *string `json:"certificate_value,omitempty"`

	// HTTPS协议使用的私钥，当证书开启时必传。取值范围：PEM编码格式。
	PrivateKey *string `json:"private_key,omitempty"`

	// 证书来源。1：代表华为云托管证书；0：表示自有证书。 默认值0。当证书开启时必传。
	CertificateSource *int32 `json:"certificate_source,omitempty"`

	// 是否使用HTTP2.0。（on：是，off：否。）,默认关闭，https_status=off时，该值不生效。
	Http2Status *string `json:"http2_status,omitempty"`
}

func (o HttpPutBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpPutBody struct{}"
	}

	return strings.Join([]string{"HttpPutBody", string(data)}, " ")
}
