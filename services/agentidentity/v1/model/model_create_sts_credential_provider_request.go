package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStsCredentialProviderRequest Request Object
type CreateStsCredentialProviderRequest struct {
	Body *CreateStsCredentialProviderReqBody `json:"body,omitempty"`
}

func (o CreateStsCredentialProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStsCredentialProviderRequest struct{}"
	}

	return strings.Join([]string{"CreateStsCredentialProviderRequest", string(data)}, " ")
}
