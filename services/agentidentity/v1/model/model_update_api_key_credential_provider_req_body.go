package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateApiKeyCredentialProviderReqBody struct {

	// The updated API key for the credential provider.
	ApiKey *string `json:"api_key,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdateApiKeyCredentialProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiKeyCredentialProviderReqBody struct{}"
	}

	return strings.Join([]string{"UpdateApiKeyCredentialProviderReqBody", string(data)}, " ")
}
