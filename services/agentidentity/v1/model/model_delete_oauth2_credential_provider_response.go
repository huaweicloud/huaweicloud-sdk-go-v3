package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOauth2CredentialProviderResponse Response Object
type DeleteOauth2CredentialProviderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOauth2CredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOauth2CredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"DeleteOauth2CredentialProviderResponse", string(data)}, " ")
}
