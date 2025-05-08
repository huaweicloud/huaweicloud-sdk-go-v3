package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResendReq struct {

	// Group ID。
	Group string `json:"group"`

	// 消息所属Topic。
	Topic *string `json:"topic,omitempty"`

	// 客户端ID。
	ClientId string `json:"client_id"`

	// 消息列表。
	MsgIdList *[]string `json:"msg_id_list,omitempty"`
}

func (o ResendReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResendReq struct{}"
	}

	return strings.Join([]string{"ResendReq", string(data)}, " ")
}
