package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermRulesResponse Response Object
type ListPermRulesResponse struct {

	// 权限信息
	Rules          *[]OnePermRuleResponseInfo `json:"rules,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListPermRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermRulesResponse struct{}"
	}

	return strings.Join([]string{"ListPermRulesResponse", string(data)}, " ")
}
