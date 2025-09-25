package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChatRsp struct {

	// **参数解释**： 对话ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 对话名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Title *string `json:"title,omitempty"`

	// **参数解释**： 对话别名。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Alias *string `json:"alias,omitempty"`

	// **参数解释**： 对话创建者。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Creator *string `json:"creator,omitempty"`

	// **参数解释**： 创建对话时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 更新对话时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 置顶对话时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TopTime *string `json:"top_time,omitempty"`
}

func (o ChatRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatRsp struct{}"
	}

	return strings.Join([]string{"ChatRsp", string(data)}, " ")
}
