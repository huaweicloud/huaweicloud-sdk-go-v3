package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateChildDto struct {

	// **参数解释：**  子节点实例ID列表。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	ChildList []string `json:"childList"`

	// **参数解释：**  父节点实例ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	ParentId string `json:"parentId"`
}

func (o BatchOperateChildDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateChildDto struct{}"
	}

	return strings.Join([]string{"BatchOperateChildDto", string(data)}, " ")
}
