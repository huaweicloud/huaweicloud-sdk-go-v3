package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRuleDetailResponse struct {
	Rule           *RuleResponse `json:"rule,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowRuleDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowRuleDetailResponse", string(data)}, " ")
}
