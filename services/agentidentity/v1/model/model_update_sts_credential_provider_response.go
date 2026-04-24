package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStsCredentialProviderResponse Response Object
type UpdateStsCredentialProviderResponse struct {
	CredentialProvider *StsCredentialProvider `json:"credential_provider,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o UpdateStsCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStsCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"UpdateStsCredentialProviderResponse", string(data)}, " ")
}
