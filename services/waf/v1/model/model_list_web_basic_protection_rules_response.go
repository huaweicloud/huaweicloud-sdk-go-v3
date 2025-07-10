package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebBasicProtectionRulesResponse Response Object
type ListWebBasicProtectionRulesResponse struct {

	// 该策略下web基础防护内置规则数量
	Total *int32 `json:"total,omitempty"`

	// web基础防护内置规则数组
	Items          *[]WebBasicProtectionRulesItem `json:"items,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListWebBasicProtectionRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebBasicProtectionRulesResponse struct{}"
	}

	return strings.Join([]string{"ListWebBasicProtectionRulesResponse", string(data)}, " ")
}
