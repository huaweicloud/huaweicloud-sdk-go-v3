package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAdvancedIpsRulesResponse Response Object
type ListAdvancedIpsRulesResponse struct {
	Data           *AdvancedIpsRuleListVo `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAdvancedIpsRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAdvancedIpsRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAdvancedIpsRulesResponse", string(data)}, " ")
}
