package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApiKeyCredentialProviderResponse Response Object
type CreateApiKeyCredentialProviderResponse struct {
	CredentialProvider *ApiKeyCredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o CreateApiKeyCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiKeyCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"CreateApiKeyCredentialProviderResponse", string(data)}, " ")
}
