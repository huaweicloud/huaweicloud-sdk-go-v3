package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTransferIpdWorkItemFlowResponse Response Object
type BatchTransferIpdWorkItemFlowResponse struct {

	// **参数解释**： 状态码。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	Result         *BatchResultVoIssueWithReasonVo `json:"result,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o BatchTransferIpdWorkItemFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTransferIpdWorkItemFlowResponse struct{}"
	}

	return strings.Join([]string{"BatchTransferIpdWorkItemFlowResponse", string(data)}, " ")
}
