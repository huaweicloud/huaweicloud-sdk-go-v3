package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodeEvent struct {

	// **参数解释**： 事件类型。 **约束限制**： 不涉及。 **取值范围**： - merge_request：MR 触发。 - push：代码Push触发。 - tag_push：标签触发。 - issue：Gitee仓库ISSUE触发。 - note：Gitee仓库评论触发。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 代码仓是否可用。 **约束限制**： 不涉及。 **取值范围**： - true：代码仓可用。 - false：代码仓不可用。 **默认取值**： 不涉及。
	Enable *bool `json:"enable,omitempty"`
}

func (o CodeEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeEvent struct{}"
	}

	return strings.Join([]string{"CodeEvent", string(data)}, " ")
}
