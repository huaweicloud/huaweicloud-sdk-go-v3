package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoliciesRequest Request Object
type ListPoliciesRequest struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListPoliciesRequest", string(data)}, " ")
}
