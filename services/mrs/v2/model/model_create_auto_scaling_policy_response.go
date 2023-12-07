package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutoScalingPolicyResponse Response Object
type CreateAutoScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAutoScalingPolicyResponse", string(data)}, " ")
}
