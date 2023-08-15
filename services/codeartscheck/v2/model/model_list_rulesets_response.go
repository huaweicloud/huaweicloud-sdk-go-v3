package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRulesetsResponse Response Object
type ListRulesetsResponse struct {

	// 规则集列表信息
	Info *[]RulesetItem `json:"info,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRulesetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRulesetsResponse struct{}"
	}

	return strings.Join([]string{"ListRulesetsResponse", string(data)}, " ")
}
