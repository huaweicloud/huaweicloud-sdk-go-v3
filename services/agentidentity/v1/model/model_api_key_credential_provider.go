package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiKeyCredentialProvider struct {

	// The name of the credential provider.
	Name string `json:"name"`

	// The Uniform Resource Name (URN) of the credential provider.
	Urn string `json:"urn"`

	ApiKeySecret *Secret `json:"api_key_secret"`

	// Timestamp in RFC 3339 format (UTC)
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o ApiKeyCredentialProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiKeyCredentialProvider struct{}"
	}

	return strings.Join([]string{"ApiKeyCredentialProvider", string(data)}, " ")
}
