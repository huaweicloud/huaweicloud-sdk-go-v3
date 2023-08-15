package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowL7policyRequest Request Object
type ShowL7policyRequest struct {

	// 转发策略id
	L7policyId string `json:"l7policy_id"`
}

func (o ShowL7policyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7policyRequest struct{}"
	}

	return strings.Join([]string{"ShowL7policyRequest", string(data)}, " ")
}
