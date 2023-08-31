package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTuningRequest Request Object
type ShowTuningRequest struct {

	// 诊断messageId
	MessageId string `json:"message_id"`

	// 连接Id
	ConnectionId string `json:"connection_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowTuningRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTuningRequest struct{}"
	}

	return strings.Join([]string{"ShowTuningRequest", string(data)}, " ")
}
