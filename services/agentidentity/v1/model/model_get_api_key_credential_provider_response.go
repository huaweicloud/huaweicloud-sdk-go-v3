package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetApiKeyCredentialProviderResponse Response Object
type GetApiKeyCredentialProviderResponse struct {
	CredentialProvider *ApiKeyCredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o GetApiKeyCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetApiKeyCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"GetApiKeyCredentialProviderResponse", string(data)}, " ")
}
