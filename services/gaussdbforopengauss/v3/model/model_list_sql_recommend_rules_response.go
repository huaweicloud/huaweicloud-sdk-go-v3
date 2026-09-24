package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlRecommendRulesResponse Response Object
type ListSqlRecommendRulesResponse struct {

	// **参数解释**: 推荐规则列表。
	RecommendRules *[]ListSqlRecommendRulesResponseResult `json:"recommend_rules,omitempty"`

	// **参数解释**: 推荐总数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlRecommendRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlRecommendRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlRecommendRulesResponse", string(data)}, " ")
}
