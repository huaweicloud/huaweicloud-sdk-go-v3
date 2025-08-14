package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCertificateResponse Response Object
type AddCertificateResponse struct {

	// CA证书ID，在上传CA证书时由平台分配的唯一标识。
	CertificateId *string `json:"certificate_id,omitempty"`

	// CA证书CN名称。
	CnName *string `json:"cn_name,omitempty"`

	// CA证书所有者。
	Owner *string `json:"owner,omitempty"`

	// CA证书验证状态。true代表证书已通过验证，可进行设备证书认证接入。false代表证书未通过验证。
	Status *bool `json:"status,omitempty"`

	// CA证书验证码。
	VerifyCode *string `json:"verify_code,omitempty"`

	// 是否开启自注册能力，当为true时该功能必须配合自注册模板使用，true：是，false：否。
	ProvisionEnable *bool `json:"provision_enable,omitempty"`

	// 绑定的自注册模板ID。
	TemplateId *string `json:"template_id,omitempty"`

	// 是否开启该CA证书下的设备证书OCSP校验，当为true且设备证书信息中包含OCSP url时则平台会校验证书的状态，当证书状态为revoked时平台会拒绝设备连接，true：开启，false：关闭。
	OcspEnable *bool `json:"ocsp_enable,omitempty"`

	// ocsp服务器端CA证书id，仅当ocsp服务器开启SSL时配置，平台使用该CA证书认证ocsp服务器。
	OcspServerCaId *string `json:"ocsp_server_ca_id,omitempty"`

	// ocsp服务器是否开启SSL加密，开启后必须配置OCSP服务器CA证书。
	OcspSslEnable *bool `json:"ocsp_ssl_enable,omitempty"`

	// 创建证书日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	CreateDate *string `json:"create_date,omitempty"`

	// CA证书生效日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	EffectiveDate *string `json:"effective_date,omitempty"`

	// CA证书失效日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	ExpiryDate     *string `json:"expiry_date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCertificateResponse struct{}"
	}

	return strings.Join([]string{"AddCertificateResponse", string(data)}, " ")
}
