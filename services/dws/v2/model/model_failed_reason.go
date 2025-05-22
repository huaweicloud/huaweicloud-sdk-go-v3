package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailedReason **参数解释**： 失败原因。如果为空，则集群处于正常状态。当存在失败的任务信息时这里可能会用来描述具体失败原因。 **取值范围**： 不涉及。
type FailedReason struct {

	// **参数解释**： 错误码。 **取值范围**： 不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o FailedReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedReason struct{}"
	}

	return strings.Join([]string{"FailedReason", string(data)}, " ")
}
