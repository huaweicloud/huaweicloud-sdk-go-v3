package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyName Human-readable display name for the policy
type PolicyName struct {
}

func (o PolicyName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyName struct{}"
	}

	return strings.Join([]string{"PolicyName", string(data)}, " ")
}
