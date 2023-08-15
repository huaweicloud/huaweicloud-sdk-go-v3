package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyRequest Request Object
type DeletePolicyRequest struct {

	// 策略的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequest", string(data)}, " ")
}
