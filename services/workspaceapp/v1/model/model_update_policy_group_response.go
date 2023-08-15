package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyGroupResponse Response Object
type UpdatePolicyGroupResponse struct {

	// 被修改策略主键
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyGroupResponse", string(data)}, " ")
}
