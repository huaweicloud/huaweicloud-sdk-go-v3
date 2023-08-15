package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeScalingPolicyResponse Response Object
type ResumeScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumeScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ResumeScalingPolicyResponse", string(data)}, " ")
}
