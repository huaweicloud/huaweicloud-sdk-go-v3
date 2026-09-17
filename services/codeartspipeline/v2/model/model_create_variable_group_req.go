package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVariableGroupReq **参数解释**： 请求体详情。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type CreateVariableGroupReq struct {

	// **参数解释**： 参数组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 描述详情。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 参数列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Variables *[]CreateVariableGroupReqVariables `json:"variables,omitempty"`
}

func (o CreateVariableGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableGroupReq struct{}"
	}

	return strings.Join([]string{"CreateVariableGroupReq", string(data)}, " ")
}
