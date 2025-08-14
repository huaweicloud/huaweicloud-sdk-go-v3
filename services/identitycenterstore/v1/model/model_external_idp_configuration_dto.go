package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExternalIdpConfigurationDto struct {

	// 身份提供商证书对应的全局唯一标识符列表
	IdpCertificateIds []IdpCertificateBody `json:"idp_certificate_ids"`

	// 身份提供商对应的全局唯一标识符（ID）
	IdpId string `json:"idp_id"`

	IdpSamlConfig *IdpSamlConfig `json:"idp_saml_config"`

	// 是否启用身份提供商
	IsEnabled bool `json:"is_enabled"`
}

func (o ExternalIdpConfigurationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalIdpConfigurationDto struct{}"
	}

	return strings.Join([]string{"ExternalIdpConfigurationDto", string(data)}, " ")
}
