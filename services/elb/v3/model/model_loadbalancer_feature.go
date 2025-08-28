package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadbalancerFeature **参数解释**：ELB实例特性。
type LoadbalancerFeature struct {

	// **参数解释**：特性名称。  **取值范围**：不涉及
	Feature string `json:"feature"`

	// **参数解释**：特性值(value字段)的类型，如：INT，表示整型。  **取值范围**：不涉及
	Type string `json:"type"`

	// **参数解释**：特性值。如开关类型的特性取值true/false，表示特性开启关闭；配额类型的特性取值整数，表示限制配额。  **取值范围**：不涉及
	Value string `json:"value"`
}

func (o LoadbalancerFeature) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadbalancerFeature struct{}"
	}

	return strings.Join([]string{"LoadbalancerFeature", string(data)}, " ")
}
