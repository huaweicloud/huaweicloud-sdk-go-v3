package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntileakagePolicyRulesResponse Response Object
type ListAntileakagePolicyRulesResponse struct {

	// 防泄露规则数量
	Total *int32 `json:"total,omitempty"`

	// 防泄露规则列表
	Items          *[]LeakageListInfo `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAntileakagePolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntileakagePolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAntileakagePolicyRulesResponse", string(data)}, " ")
}
