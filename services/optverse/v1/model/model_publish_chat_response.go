package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishChatResponse Response Object
type PublishChatResponse struct {

	// **参数解释**： 发布ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishChatResponse struct{}"
	}

	return strings.Join([]string{"PublishChatResponse", string(data)}, " ")
}
