package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateChildDto struct {

	// **参数解释：**  子节点实例ID列表，用于指定待批量添加为目标父节点子节点的实例。列表中的每个元素对应一个需要挂载的子节点实例ID，这些实例当前必须没有父节点。  **约束限制：**  子节点实例当前必须没有父节点。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	ChildList []string `json:"childList"`

	// **参数解释：**  父节点实例ID，用于指定待添加子节点所属的目标父节点。例如BOM中的某个装配节点ID、组织架构中的某个部门ID等。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	ParentId string `json:"parentId"`
}

func (o BatchOperateChildDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateChildDto struct{}"
	}

	return strings.Join([]string{"BatchOperateChildDto", string(data)}, " ")
}
