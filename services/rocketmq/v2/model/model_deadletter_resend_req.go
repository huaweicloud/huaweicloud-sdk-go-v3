package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeadletterResendReq struct {

	// Topic名称。
	Topic *string `json:"topic,omitempty"`

	// 消息列表。
	MsgIdList *[]string `json:"msg_id_list,omitempty"`
}

func (o DeadletterResendReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadletterResendReq struct{}"
	}

	return strings.Join([]string{"DeadletterResendReq", string(data)}, " ")
}
