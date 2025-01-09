package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTargetOfPolicyGroupResponse Response Object
type ListTargetOfPolicyGroupResponse struct {

	// 应用对象列表。
	Targets *[]Target `json:"targets,omitempty"`

	// 用户列表数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTargetOfPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTargetOfPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListTargetOfPolicyGroupResponse", string(data)}, " ")
}
