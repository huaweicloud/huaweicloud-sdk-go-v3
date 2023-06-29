package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTargetsOfPolicyGroupResponse Response Object
type ListTargetsOfPolicyGroupResponse struct {

	// 应用对象列表。
	Targets        *[]Target `json:"targets,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTargetsOfPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTargetsOfPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListTargetsOfPolicyGroupResponse", string(data)}, " ")
}
