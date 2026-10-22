package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReferencedPolicy struct {

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`
}

func (o ReferencedPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReferencedPolicy struct{}"
	}

	return strings.Join([]string{"ReferencedPolicy", string(data)}, " ")
}
