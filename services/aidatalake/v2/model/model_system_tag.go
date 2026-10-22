package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemTag **参数解释**：系统标签。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type SystemTag struct {

	// **参数解释**：系统标签键。 **约束限制**：不涉及。 **取值范围**：目前仅支持 _sys_enterprise_project_id。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：系统标签值。企业项目ID，获取方法，请参见[获取企业项目ID](aidatalake_02_0042.xml)。 **约束限制**：不涉及。 **取值范围**：长度为1~64个字符，由英文字母、数字和中划线组成。 **默认取值**：不涉及。
	Value string `json:"value"`
}

func (o SystemTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemTag struct{}"
	}

	return strings.Join([]string{"SystemTag", string(data)}, " ")
}
