package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpPolicyRequest Request Object
type DeleteHttpPolicyRequest struct {

	// 防护策略id
	PolicyId string `json:"policy_id"`
}

func (o DeleteHttpPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpPolicyRequest", string(data)}, " ")
}
