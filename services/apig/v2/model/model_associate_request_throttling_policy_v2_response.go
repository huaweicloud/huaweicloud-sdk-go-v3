package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateRequestThrottlingPolicyV2Response struct {
	// API与流控策略的绑定关系列表

	ThrottleApplys *[]ThrottleBindingResp `json:"throttle_applys,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o AssociateRequestThrottlingPolicyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateRequestThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"AssociateRequestThrottlingPolicyV2Response", string(data)}, " ")
}
