package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskConfigCheckRulesResponse Response Object
type ListRiskConfigCheckRulesResponse struct {

	// **参数解释**: 风险总数 **取值范围**: 不涉及
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 数据列表 **取值范围**: 不涉及
	DataList       *[]CheckRuleRiskInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListRiskConfigCheckRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigCheckRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRiskConfigCheckRulesResponse", string(data)}, " ")
}
