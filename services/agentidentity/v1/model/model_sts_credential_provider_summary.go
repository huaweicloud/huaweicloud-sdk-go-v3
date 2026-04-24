package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StsCredentialProviderSummary struct {

	// The name of the credential provider.
	Name string `json:"name"`

	// 凭证提供者的唯一资源名称（URN）。
	Urn string `json:"urn"`

	// The URN of the agency used to obtain IAM temporary credentials.
	AgencyUrn string `json:"agency_urn"`

	// Timestamp in RFC 3339 format (UTC)
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o StsCredentialProviderSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StsCredentialProviderSummary struct{}"
	}

	return strings.Join([]string{"StsCredentialProviderSummary", string(data)}, " ")
}
