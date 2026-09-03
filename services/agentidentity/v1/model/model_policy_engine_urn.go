package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyEngineUrn The URN of the policy engine.
type PolicyEngineUrn struct {
}

func (o PolicyEngineUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEngineUrn struct{}"
	}

	return strings.Join([]string{"PolicyEngineUrn", string(data)}, " ")
}
