package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListScalingPoliciesResponse struct {

	// 总记录数。
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number"`

	// 查询的起始行号。
	StartNumber *int32 `json:"start_number,omitempty" xml:"start_number"`

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 伸缩策略列表
	ScalingPolicies *[]ScalingV1PolicyDetail `json:"scaling_policies,omitempty" xml:"scaling_policies"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ListScalingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListScalingPoliciesResponse", string(data)}, " ")
}
