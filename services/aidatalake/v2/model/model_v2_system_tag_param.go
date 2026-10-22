package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2SystemTagParam **参数解释**：系统标签查询参数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type V2SystemTagParam struct {

	// **参数解释**：系统标签的键。 **约束限制**：不涉及。 **取值范围**：目前仅支持 _sys_enterprise_project_id。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：系统标签的值列表，即企业项目ID列表。企业项目ID可通过[获取企业项目ID](aidatalake_02_0042.xml)获取。 **约束限制**：不涉及。
	Values []string `json:"values"`
}

func (o V2SystemTagParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2SystemTagParam struct{}"
	}

	return strings.Join([]string{"V2SystemTagParam", string(data)}, " ")
}
