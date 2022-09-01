package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRulesetsResponse struct {

	// 规则集列表信息
	Info *[]RulesetItem `json:"info,omitempty" xml:"info"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRulesetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRulesetsResponse struct{}"
	}

	return strings.Join([]string{"ListRulesetsResponse", string(data)}, " ")
}
