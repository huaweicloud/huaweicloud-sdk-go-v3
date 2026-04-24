package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkloadIdentity struct {

	// The name of the workload identity.
	Name string `json:"name"`

	// The URN of the workload identity.
	Urn string `json:"urn"`

	AllowedResourceOauth2ReturnUrls *[]string `json:"allowed_resource_oauth2_return_urls,omitempty"`

	AuthorizerType *AuthorizerType `json:"authorizer_type"`

	ApiKeySecret *Secret `json:"api_key_secret,omitempty"`

	CreatedBy *CreatedBy `json:"created_by"`

	// Timestamp in RFC 3339 format (UTC)
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o WorkloadIdentity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadIdentity struct{}"
	}

	return strings.Join([]string{"WorkloadIdentity", string(data)}, " ")
}
