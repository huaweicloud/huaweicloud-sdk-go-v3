package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomPolicyRulesResponse Response Object
type ListCustomPolicyRulesResponse struct {

	// 数量
	Total *int32 `json:"total,omitempty"`

	// 防护规则列表
	Items          *[]CustomRule `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCustomPolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomPolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListCustomPolicyRulesResponse", string(data)}, " ")
}
