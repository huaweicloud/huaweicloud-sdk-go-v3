package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IefFlinkJobMessagesReq IEF Flink job action回调的请求body体。
type IefFlinkJobMessagesReq struct {

	// 消息id
	MessageId string `json:"message_id"`

	State *State `json:"state,omitempty"`
}

func (o IefFlinkJobMessagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IefFlinkJobMessagesReq struct{}"
	}

	return strings.Join([]string{"IefFlinkJobMessagesReq", string(data)}, " ")
}
