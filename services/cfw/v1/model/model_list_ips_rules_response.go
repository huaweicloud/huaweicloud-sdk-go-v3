package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsRulesResponse Response Object
type ListIpsRulesResponse struct {
	Data           *AdvancedIpsRuleListVo `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListIpsRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsRulesResponse struct{}"
	}

	return strings.Join([]string{"ListIpsRulesResponse", string(data)}, " ")
}
