package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessScalingPolicyResponse Response Object
type UpdateServerlessScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerlessScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerlessScalingPolicyResponse", string(data)}, " ")
}
