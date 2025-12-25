package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrashcanResultData struct {

	// **参数解释**： 成功数目。 **取值范围**： 不涉及。
	SuccessNum *int32 `json:"successNum,omitempty"`

	// **参数解释**： 失败数目。 **取值范围**： 不涉及。
	FailedNum *int32 `json:"failedNum,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailedMessages *[]string `json:"failedMessages,omitempty"`
}

func (o TrashcanResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrashcanResultData struct{}"
	}

	return strings.Join([]string{"TrashcanResultData", string(data)}, " ")
}
