package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyUrn The URN of the policy.
type PolicyUrn struct {
}

func (o PolicyUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyUrn struct{}"
	}

	return strings.Join([]string{"PolicyUrn", string(data)}, " ")
}
