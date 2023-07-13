package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyGroupResponse Response Object
type CreatePolicyGroupResponse struct {

	// 策略组ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyGroupResponse", string(data)}, " ")
}
