package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSendDiffSmsRequest Request Object
type BatchSendDiffSmsRequest struct {
	Body *BatchSendDiffSmsRequestBody `json:"body,omitempty"`
}

func (o BatchSendDiffSmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendDiffSmsRequest struct{}"
	}

	return strings.Join([]string{"BatchSendDiffSmsRequest", string(data)}, " ")
}
