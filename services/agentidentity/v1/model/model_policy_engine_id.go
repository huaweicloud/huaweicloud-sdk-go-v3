package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyEngineId System-generated unique identifier for the policy engine.
type PolicyEngineId struct {
}

func (o PolicyEngineId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEngineId struct{}"
	}

	return strings.Join([]string{"PolicyEngineId", string(data)}, " ")
}
