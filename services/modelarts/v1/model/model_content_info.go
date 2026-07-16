package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContentInfo 算法内容详情
type ContentInfo struct {

	// **参数解释**：版本数量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	VersionNum *string `json:"version_num,omitempty"`

	// **参数解释**：描述。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Desc *string `json:"desc,omitempty"`
}

func (o ContentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentInfo struct{}"
	}

	return strings.Join([]string{"ContentInfo", string(data)}, " ")
}
