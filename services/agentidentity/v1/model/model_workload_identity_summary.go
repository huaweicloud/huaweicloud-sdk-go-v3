package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkloadIdentitySummary struct {

	// The name of the workload identity.
	Name string `json:"name"`

	// The URN of the workload identity.
	Urn string `json:"urn"`

	AuthorizerType *AuthorizerType `json:"authorizer_type"`

	CreatedBy *CreatedBy `json:"created_by"`

	// Timestamp in RFC 3339 format (UTC)
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o WorkloadIdentitySummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadIdentitySummary struct{}"
	}

	return strings.Join([]string{"WorkloadIdentitySummary", string(data)}, " ")
}
