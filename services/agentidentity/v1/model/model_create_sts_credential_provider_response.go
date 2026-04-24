package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStsCredentialProviderResponse Response Object
type CreateStsCredentialProviderResponse struct {
	CredentialProvider *StsCredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o CreateStsCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStsCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"CreateStsCredentialProviderResponse", string(data)}, " ")
}
