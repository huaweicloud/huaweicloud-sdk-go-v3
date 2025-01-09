package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyGroupInfoResponse Response Object
type ListPolicyGroupInfoResponse struct {

	// 策略组。
	PolicyGroups *[]PolicyGroupForList `json:"policy_groups,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPolicyGroupInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyGroupInfoResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyGroupInfoResponse", string(data)}, " ")
}
