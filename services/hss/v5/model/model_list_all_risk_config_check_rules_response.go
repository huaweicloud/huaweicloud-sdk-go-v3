package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllRiskConfigCheckRulesResponse Response Object
type ListAllRiskConfigCheckRulesResponse struct {

	// **参数解释** 风险总数 **取值范围** 取值0-200000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 已通过检查项数量 **取值范围** 取值0-200000
	PassNum *int32 `json:"pass_num,omitempty"`

	// **参数解释** 未通过检查项数量 **取值范围** 取值0-200000
	FailedNum *int32 `json:"failed_num,omitempty"`

	// **参数解释** 已处理检查项数量 **取值范围** 取值0-200000
	ProcessedNum *int32 `json:"processed_num,omitempty"`

	// **参数解释** 数据列表
	DataList       *[]CheckRulesResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAllRiskConfigCheckRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllRiskConfigCheckRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAllRiskConfigCheckRulesResponse", string(data)}, " ")
}
