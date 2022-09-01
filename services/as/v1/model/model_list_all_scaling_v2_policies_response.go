package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAllScalingV2PoliciesResponse struct {

	// 总记录数。
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number"`

	// 查询的起始行号。
	StartNumber *int32 `json:"start_number,omitempty" xml:"start_number"`

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 伸缩策略列表
	ScalingPolicies *[]ScalingAllPolicyDetail `json:"scaling_policies,omitempty" xml:"scaling_policies"`
	HttpStatusCode  int                       `json:"-"`
}

func (o ListAllScalingV2PoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllScalingV2PoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListAllScalingV2PoliciesResponse", string(data)}, " ")
}
