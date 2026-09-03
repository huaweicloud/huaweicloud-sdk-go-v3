package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyEngineName Customer-assigned immutable name for the policy engine.
type PolicyEngineName struct {
}

func (o PolicyEngineName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEngineName struct{}"
	}

	return strings.Join([]string{"PolicyEngineName", string(data)}, " ")
}
