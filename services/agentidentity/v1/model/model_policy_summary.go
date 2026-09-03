package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicySummary struct {

	// System-generated unique identifier for the policy.
	PolicyId string `json:"policy_id"`

	// Human-readable display name for the policy
	Name string `json:"name"`

	// 策略的可读描述。
	Description *string `json:"description,omitempty"`

	// The URN of the policy.
	Urn string `json:"urn"`

	Status *PolicyStatus `json:"status"`

	// 关于策略状态的额外信息，提供关于任何失败或策略创建过程当前状态的详细信息。
	StatusReasons *[]string `json:"status_reasons,omitempty"`

	// Timestamp in RFC 3339 format (UTC)
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`
}

func (o PolicySummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicySummary struct{}"
	}

	return strings.Join([]string{"PolicySummary", string(data)}, " ")
}
