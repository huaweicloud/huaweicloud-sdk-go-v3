package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRuleHitCountRequest struct {
	Body *ListRuleHitCountDto `json:"body,omitempty"`
}

func (o ListRuleHitCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleHitCountRequest struct{}"
	}

	return strings.Join([]string{"ListRuleHitCountRequest", string(data)}, " ")
}
