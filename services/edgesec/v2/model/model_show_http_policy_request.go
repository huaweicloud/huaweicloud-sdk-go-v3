package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpPolicyRequest Request Object
type ShowHttpPolicyRequest struct {

	// 防护策略id
	PolicyId string `json:"policy_id"`
}

func (o ShowHttpPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpPolicyRequest", string(data)}, " ")
}
