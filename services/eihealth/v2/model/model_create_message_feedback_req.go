package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMessageFeedbackReq struct {
	Type *FeedbackType `json:"type"`

	// **参数解释**： 反馈标签。 **约束限制**： 最多支持设置10个标签。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Labels *[]FeedbackLabel `json:"labels,omitempty"`

	// **参数解释**： 反馈内容。 **约束限制**： 不涉及 **取值范围**： 长度为[0-1000]个字符。 **默认取值**： 不涉及
	Content *string `json:"content,omitempty"`
}

func (o CreateMessageFeedbackReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageFeedbackReq struct{}"
	}

	return strings.Join([]string{"CreateMessageFeedbackReq", string(data)}, " ")
}
