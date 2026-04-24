package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetStsCredentialProviderResponse Response Object
type GetStsCredentialProviderResponse struct {
	CredentialProvider *StsCredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o GetStsCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetStsCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"GetStsCredentialProviderResponse", string(data)}, " ")
}
