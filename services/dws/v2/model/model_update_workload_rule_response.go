package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkloadRuleResponse Response Object
type UpdateWorkloadRuleResponse struct {

	// **参数解释**： 错误码，0表示成功。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 错误信息，成功时为空。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateWorkloadRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkloadRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkloadRuleResponse", string(data)}, " ")
}
