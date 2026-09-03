package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyEngineSummary struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`

	// Customer-assigned immutable name for the policy engine.
	Name string `json:"name"`

	Type *PolicyEngineType `json:"type"`

	// The URN of the policy engine.
	Urn string `json:"urn"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// Timestamp in RFC 3339 format (UTC)
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o PolicyEngineSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEngineSummary struct{}"
	}

	return strings.Join([]string{"PolicyEngineSummary", string(data)}, " ")
}
