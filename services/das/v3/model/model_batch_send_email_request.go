package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSendEmailRequest Request Object
type BatchSendEmailRequest struct {
	Body *BatchSendEmailRequestBody `json:"body,omitempty"`
}

func (o BatchSendEmailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendEmailRequest struct{}"
	}

	return strings.Join([]string{"BatchSendEmailRequest", string(data)}, " ")
}
