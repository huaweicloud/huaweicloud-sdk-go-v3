package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetadataTask **参数解释**： 元数据迁移任务。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type MetadataTask struct {

	// **参数解释**： 元数据迁移任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 元数据迁移任务名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 元数据迁移任务开始时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartDate *string `json:"start_date,omitempty"`

	// **参数解释**： 元数据迁移任务状态。 **约束限制**： 不涉及。 **取值范围**： - creating：创建中。 - starting：迁移中。 - failed：迁移失败。 - finished：迁移完成。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 元数据迁移类型。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：从rocketmq迁移到rocketmq。 - rabbitToRocket：从rabbitmq迁移到rocketmq。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o MetadataTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataTask struct{}"
	}

	return strings.Join([]string{"MetadataTask", string(data)}, " ")
}
