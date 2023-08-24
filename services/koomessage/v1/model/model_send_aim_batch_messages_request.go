package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendAimBatchMessagesRequest Request Object
type SendAimBatchMessagesRequest struct {
	Body *SmsTaskReq `json:"body,omitempty"`
}

func (o SendAimBatchMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendAimBatchMessagesRequest struct{}"
	}

	return strings.Join([]string{"SendAimBatchMessagesRequest", string(data)}, " ")
}
