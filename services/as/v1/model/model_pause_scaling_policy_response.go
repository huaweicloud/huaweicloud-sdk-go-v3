package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PauseScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"PauseScalingPolicyResponse", string(data)}, " ")
}
