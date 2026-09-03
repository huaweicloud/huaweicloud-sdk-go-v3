package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyDefinition Represents the definition structure for policies. This structure encapsulates different policy formats and languages that can be used to define access control rules. This is a UNION, so only one of the following members can be specified when used or returned.
type PolicyDefinition struct {
	Cedar *CedarPolicy `json:"cedar"`
}

func (o PolicyDefinition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDefinition struct{}"
	}

	return strings.Join([]string{"PolicyDefinition", string(data)}, " ")
}
