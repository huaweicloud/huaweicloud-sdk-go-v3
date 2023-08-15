package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenApiParaForCheckMessage struct {

	// 消息编号
	MessageId *string `json:"message_id,omitempty"`

	// 执行动作。0=立刻执行, 1=定期执行。
	Action *int32 `json:"action,omitempty"`

	// 使用截止时间。仅定期执行需要此参数，默认服务器当前时间三天后。
	Time *string `json:"time,omitempty"`
}

func (o OpenApiParaForCheckMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApiParaForCheckMessage struct{}"
	}

	return strings.Join([]string{"OpenApiParaForCheckMessage", string(data)}, " ")
}
