package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSendSmsRequest Request Object
type BatchSendSmsRequest struct {
	Body *BatchSendSmsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o BatchSendSmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendSmsRequest struct{}"
	}

	return strings.Join([]string{"BatchSendSmsRequest", string(data)}, " ")
}
