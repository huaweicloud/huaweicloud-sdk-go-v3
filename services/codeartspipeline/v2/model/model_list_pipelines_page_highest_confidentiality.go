package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelinesPageHighestConfidentiality **参数解释**： 最高密级。 **约束限制**： 非涉密场景无该字段。 **取值范围**： 不涉及。
type ListPipelinesPageHighestConfidentiality struct {

	// **参数解释**： 密级ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 密级等级逻辑ID。 **取值范围**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 预留字段。 **取值范围**： 不涉及。
	Reserve1 *string `json:"reserve_1,omitempty"`

	// **参数解释**： 密级等级中文名。 **取值范围**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 密级等级英文名。 **取值范围**： 不涉及。
	ValueEn *string `json:"value_en,omitempty"`

	// **参数解释**： 密级等级序号，密级越高数字越大。 **取值范围**： 正整数。
	Sequence *int32 `json:"sequence,omitempty"`
}

func (o ListPipelinesPageHighestConfidentiality) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesPageHighestConfidentiality struct{}"
	}

	return strings.Join([]string{"ListPipelinesPageHighestConfidentiality", string(data)}, " ")
}
