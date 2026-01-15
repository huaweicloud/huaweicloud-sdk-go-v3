package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiOpsRiskObject struct {

	// **参数解释**： 检测项ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 检测项名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 检测项说明 **取值范围**： 不涉及
	Desc *string `json:"desc,omitempty"`

	// **参数解释**： 检测结果风险等级 **取值范围**： 不涉及
	Level *string `json:"level,omitempty"`

	// **参数解释**： 检测结果 **取值范围**： 不涉及
	Result *string `json:"result,omitempty"`

	// **参数解释**： 检测建议 **取值范围**： 不涉及
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o AiOpsRiskObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiOpsRiskObject struct{}"
	}

	return strings.Join([]string{"AiOpsRiskObject", string(data)}, " ")
}
