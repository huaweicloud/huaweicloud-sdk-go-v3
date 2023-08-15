package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResendReq struct {

	// Group ID。
	Group *string `json:"group,omitempty"`

	// 客户端ID。
	ClientId *string `json:"client_id,omitempty"`

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
