package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOauth2CredentialProviderRequest Request Object
type CreateOauth2CredentialProviderRequest struct {
	Body *CreateOauth2CredentialProviderReqBody `json:"body,omitempty"`
}

func (o CreateOauth2CredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOauth2CredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"CreateOauth2CredentialProviderRequest", string(data)}, " ")
}
