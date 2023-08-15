package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleInfo Information of rule
type RuleInfo struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

	// Project id value
	Rule *string `json:"rule,omitempty"`
}

func (o RuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInfo struct{}"
	}

	return strings.Join([]string{"RuleInfo", string(data)}, " ")
}
