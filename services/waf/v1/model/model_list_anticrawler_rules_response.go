package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnticrawlerRulesResponse Response Object
type ListAnticrawlerRulesResponse struct {

	// 该策略下反爬虫规则数量
	Total *int32 `json:"total,omitempty"`

	// 反爬虫规则列表
	Items          *[]AnticrawlerRule `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAnticrawlerRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnticrawlerRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAnticrawlerRulesResponse", string(data)}, " ")
}
