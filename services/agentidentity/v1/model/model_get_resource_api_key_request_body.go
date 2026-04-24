package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetResourceApiKeyRequestBody struct {

	// Name of the resource credential provider to retrieve API key from
	ResourceCredentialProviderName string `json:"resource_credential_provider_name"`

	// Identity token of the workload requesting the API key
	WorkloadAccessToken *string `json:"workload_access_token,omitempty"`
}

func (o GetResourceApiKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceApiKeyRequestBody struct{}"
	}

	return strings.Join([]string{"GetResourceApiKeyRequestBody", string(data)}, " ")
}
