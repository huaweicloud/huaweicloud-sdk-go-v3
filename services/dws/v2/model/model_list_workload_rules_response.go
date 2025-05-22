package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadRulesResponse Response Object
type ListWorkloadRulesResponse struct {

	// **参数解释**： 错误编码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 查询错误详情。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	// **参数解释**： 异常规则列表。 **取值范围**： 不涉及。
	Items *[]ExceptRuleBo `json:"items,omitempty"`

	// **参数解释**： 异常规则总条数。 **取值范围**： 大于等于0
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWorkloadRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadRulesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadRulesResponse", string(data)}, " ")
}
