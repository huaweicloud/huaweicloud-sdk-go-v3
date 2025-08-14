package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalIdPConfigurationForDirectoryResponse Response Object
type CreateExternalIdPConfigurationForDirectoryResponse struct {
	HwsSpSamlConfig *SpsamlConfig `json:"hws_sp_saml_config,omitempty"`

	// 身份提供商证书全局唯一标识符（ID)
	IdpCertificateId *string `json:"idp_certificate_id,omitempty"`

	// 身份提供商证书全局唯一标识符列表
	IdpCertificateIds *[]string `json:"idp_certificate_ids,omitempty"`

	// 外部身份提供商的全局唯一标识符（ID）
	IdpId          *string `json:"idp_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateExternalIdPConfigurationForDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalIdPConfigurationForDirectoryResponse struct{}"
	}

	return strings.Join([]string{"CreateExternalIdPConfigurationForDirectoryResponse", string(data)}, " ")
}
