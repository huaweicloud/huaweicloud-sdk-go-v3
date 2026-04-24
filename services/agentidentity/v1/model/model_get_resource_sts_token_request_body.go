package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetResourceStsTokenRequestBody struct {

	// Name of the STS credential provider to retrieve STS credentials from
	ResourceCredentialProviderName string `json:"resource_credential_provider_name"`

	// Identity token of the workload requesting the STS token
	WorkloadAccessToken *string `json:"workload_access_token,omitempty"`

	// An identifier for the assumed agency session
	AgencySessionName string `json:"agency_session_name"`

	// The duration, in seconds, of the agency session
	DurationSeconds *int32 `json:"duration_seconds,omitempty"`

	// An IAM policy in JSON format that you want to use as an inline session policy
	Policy *string `json:"policy,omitempty"`

	// The source identity specified by the principal that is calling the operation
	SourceIdentity *string `json:"source_identity,omitempty"`

	// A list of tags
	Tags *[]StsTag `json:"tags,omitempty"`

	// A list of keys for session tags that you want to set as transitive
	TransitiveTagKeys *[]string `json:"transitive_tag_keys,omitempty"`
}

func (o GetResourceStsTokenRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceStsTokenRequestBody struct{}"
	}

	return strings.Join([]string{"GetResourceStsTokenRequestBody", string(data)}, " ")
}
