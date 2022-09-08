package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePolicyRequest struct {

	// 策略ID
	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequest", string(data)}, " ")
}
