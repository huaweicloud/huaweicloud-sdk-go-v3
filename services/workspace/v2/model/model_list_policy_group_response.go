package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyGroupResponse Response Object
type ListPolicyGroupResponse struct {

	// 策略组。
	PolicyGroups *[]PolicyGroupForList `json:"policy_groups,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyGroupResponse", string(data)}, " ")
}
