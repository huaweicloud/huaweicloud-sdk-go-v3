package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessageDetailRequest Request Object
type ShowMessageDetailRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 消息信息id
	MessageId string `json:"message_id"`
}

func (o ShowMessageDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMessageDetailRequest", string(data)}, " ")
}
