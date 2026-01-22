package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConsumedMessageForRocketMqResponse Response Object
type ValidateConsumedMessageForRocketMqResponse struct {

	// **参数解释**： 消费验证结果。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	ResendResults  *[]DeadletterResendRespResendResults `json:"resend_results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ValidateConsumedMessageForRocketMqResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConsumedMessageForRocketMqResponse struct{}"
	}

	return strings.Join([]string{"ValidateConsumedMessageForRocketMqResponse", string(data)}, " ")
}
