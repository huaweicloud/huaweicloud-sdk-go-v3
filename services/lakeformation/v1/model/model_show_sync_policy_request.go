package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSyncPolicyRequest Request Object
type ShowSyncPolicyRequest struct {

	// instance id
	InstanceId string `json:"instance_id"`

	// lastKnownVersion
	LastKnownVersion *int64 `json:"last_known_version,omitempty"`

	// supportsPolicyDeltas
	SupportsPolicyDeltas *bool `json:"supports_policy_deltas,omitempty"`

	// isReturnPolicyData
	IsReturnPolicyData *bool `json:"is_return_policy_data,omitempty"`
}

func (o ShowSyncPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSyncPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowSyncPolicyRequest", string(data)}, " ")
}
