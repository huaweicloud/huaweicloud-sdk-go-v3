package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateChatReq struct {

	// **参数解释**： 对话别名。 **约束限制**： 不涉及 **取值范围**： 取值范围为[1-20]个字符。 **默认取值**： 不涉及
	Alias string `json:"alias"`

	// **参数解释**： 是否置顶对话。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	EnableTop bool `json:"enable_top"`
}

func (o UpdateChatReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChatReq struct{}"
	}

	return strings.Join([]string{"UpdateChatReq", string(data)}, " ")
}
