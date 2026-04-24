package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateApiKeyCredentialProviderReqBody struct {

	// The name of the credential provider.
	Name string `json:"name"`

	// The API key used to authenticate with the external provider.
	ApiKey string `json:"api_key"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateApiKeyCredentialProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiKeyCredentialProviderReqBody struct{}"
	}

	return strings.Join([]string{"CreateApiKeyCredentialProviderReqBody", string(data)}, " ")
}
