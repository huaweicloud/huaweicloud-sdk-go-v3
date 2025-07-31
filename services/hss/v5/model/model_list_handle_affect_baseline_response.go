package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHandleAffectBaselineResponse Response Object
type ListHandleAffectBaselineResponse struct {

	// **参数解释** 该操作影响的范围的总数 **取值范围**   取值0-5000
	TotalRuleNum *int32 `json:"total_rule_num,omitempty"`

	// **参数解释** 该操作影响的检查项数 **取值范围**   取值0-5000
	RuleNum *int32 `json:"rule_num,omitempty"`

	// **参数解释** 该操作影响的主机数 **取值范围**   取值0-5000
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释** 该操作影响范围的详细信息的列表
	DataList       *[]HandleAffectBaselineInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListHandleAffectBaselineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHandleAffectBaselineResponse struct{}"
	}

	return strings.Join([]string{"ListHandleAffectBaselineResponse", string(data)}, " ")
}
