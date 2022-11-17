package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRuleHitCountResponse struct {
	Data           *RuleHitCountRecords `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListRuleHitCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleHitCountResponse struct{}"
	}

	return strings.Join([]string{"ListRuleHitCountResponse", string(data)}, " ")
}
