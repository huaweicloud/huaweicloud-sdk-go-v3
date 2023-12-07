package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutoScalingPolicyResponse Response Object
type DeleteAutoScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAutoScalingPolicyResponse", string(data)}, " ")
}
