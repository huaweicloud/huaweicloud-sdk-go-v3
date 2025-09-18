package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Workflow 流水线参数详情
type Workflow struct {

	// **参数解释**： 任务类型，list类型数据。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Parameter []PipelineParam `json:"parameter"`

	// **参数解释**： 源码仓，list类型数据。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Source []Source `json:"source"`

	// **参数解释**： 流水线名字 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 项目ID **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 项目名字 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectName string `json:"project_name"`
}

func (o Workflow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workflow struct{}"
	}

	return strings.Join([]string{"Workflow", string(data)}, " ")
}
