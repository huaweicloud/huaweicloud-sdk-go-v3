package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIdpReqBody struct {

	// 身份提供商SAML元数据，与身份提供商SAML配置二选一
	IdpSamlMetadata *string `json:"idp_saml_metadata,omitempty"`

	// 身份提供商证书，与身份提供商SAML配置联合使用
	IdpCertificate *string `json:"idp_certificate,omitempty"`

	// 身份提供商SAML配置信息，与身份提供商SAML元数据二选一
	IdpSamlConfig *interface{} `json:"idp_saml_config,omitempty"`
}

func (o CreateIdpReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIdpReqBody struct{}"
	}

	return strings.Join([]string{"CreateIdpReqBody", string(data)}, " ")
}
