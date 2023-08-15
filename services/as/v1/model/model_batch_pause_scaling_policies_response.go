package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPauseScalingPoliciesResponse Response Object
type BatchPauseScalingPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchPauseScalingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPauseScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchPauseScalingPoliciesResponse", string(data)}, " ")
}
