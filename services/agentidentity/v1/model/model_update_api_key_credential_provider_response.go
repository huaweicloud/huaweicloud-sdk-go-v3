package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApiKeyCredentialProviderResponse Response Object
type UpdateApiKeyCredentialProviderResponse struct {
	CredentialProvider *ApiKeyCredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o UpdateApiKeyCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiKeyCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"UpdateApiKeyCredentialProviderResponse", string(data)}, " ")
}
