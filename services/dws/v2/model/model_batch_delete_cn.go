package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCn **参数解释**： 批量删除CN节点ID信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type BatchDeleteCn struct {

	// **参数解释**： 批量删除CN节点ID。 **约束限制**： 不涉及。 **取值范围**： 非空。 **默认取值**： 不涉及。
	Instances *[]string `json:"instances,omitempty"`
}

func (o BatchDeleteCn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCn struct{}"
	}

	return strings.Join([]string{"BatchDeleteCn", string(data)}, " ")
}
