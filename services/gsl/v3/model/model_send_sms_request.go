package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendSmsRequest Request Object
type SendSmsRequest struct {
	Body *CreateSendSmsReq `json:"body,omitempty"`
}

func (o SendSmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendSmsRequest struct{}"
	}

	return strings.Join([]string{"SendSmsRequest", string(data)}, " ")
}
