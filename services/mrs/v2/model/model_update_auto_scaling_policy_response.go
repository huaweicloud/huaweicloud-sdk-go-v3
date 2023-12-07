package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoScalingPolicyResponse Response Object
type UpdateAutoScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoScalingPolicyResponse", string(data)}, " ")
}
