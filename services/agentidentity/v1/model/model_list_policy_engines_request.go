package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyEnginesRequest Request Object
type ListPolicyEnginesRequest struct {

	// Filter by policy engine type. If not specified, all policy engine types are returned.
	Type *PolicyEngineType `json:"type,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListPolicyEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyEnginesRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyEnginesRequest", string(data)}, " ")
}
