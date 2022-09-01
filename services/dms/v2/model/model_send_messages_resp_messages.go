package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendMessagesRespMessages struct {

	// 错误描述信息。
	Error *string `json:"error,omitempty" xml:"error"`

	// 错误码。
	ErrorCode *int32 `json:"error_code,omitempty" xml:"error_code"`

	// 发送消息的状态。 0：表示发送成功。 1：表示发送失败，失败原因参考对应的error和error_code。
	State *int32 `json:"state,omitempty" xml:"state"`

	// 消息ID。
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o SendMessagesRespMessages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessagesRespMessages struct{}"
	}

	return strings.Join([]string{"SendMessagesRespMessages", string(data)}, " ")
}
