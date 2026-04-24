package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApiKeyCredentialProviderResponse Response Object
type DeleteApiKeyCredentialProviderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiKeyCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiKeyCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"DeleteApiKeyCredentialProviderResponse", string(data)}, " ")
}
