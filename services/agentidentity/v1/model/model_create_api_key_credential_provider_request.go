package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApiKeyCredentialProviderRequest Request Object
type CreateApiKeyCredentialProviderRequest struct {
	Body *CreateApiKeyCredentialProviderReqBody `json:"body,omitempty"`
}

func (o CreateApiKeyCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiKeyCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"CreateApiKeyCredentialProviderRequest", string(data)}, " ")
}
