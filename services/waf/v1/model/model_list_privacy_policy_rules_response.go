package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivacyPolicyRulesResponse Response Object
type ListPrivacyPolicyRulesResponse struct {

	// 规则条数
	Total *int32 `json:"total,omitempty"`

	// 规则详情数组
	Items          *[]PrivacyResponseBody `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListPrivacyPolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivacyPolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivacyPolicyRulesResponse", string(data)}, " ")
}
