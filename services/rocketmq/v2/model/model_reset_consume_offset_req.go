package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetConsumeOffsetReq struct {

	// **参数解释**： 重置的主题。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 重置的时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Timestamp string `json:"timestamp"`
}

func (o ResetConsumeOffsetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetReq struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetReq", string(data)}, " ")
}
