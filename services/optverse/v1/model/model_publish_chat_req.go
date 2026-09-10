package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublishChatReq struct {

	// **参数解释**： 发布资产名称。 **约束限制**： 不涉及 **取值范围**： 取值范围为[1-128]个字符。 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 发布资产描述。 **约束限制**： 不涉及 **取值范围**： 取值范围为[0-1024]个字符。 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	Type *AgentType `json:"type"`
}

func (o PublishChatReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishChatReq struct{}"
	}

	return strings.Join([]string{"PublishChatReq", string(data)}, " ")
}
