package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRulesResponse Response Object
type ListCustomRulesResponse struct {

	// 数量
	Total *int32 `json:"total,omitempty"`

	// 防护规则列表
	Items          *[]CustomRule `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCustomRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRulesResponse struct{}"
	}

	return strings.Join([]string{"ListCustomRulesResponse", string(data)}, " ")
}
