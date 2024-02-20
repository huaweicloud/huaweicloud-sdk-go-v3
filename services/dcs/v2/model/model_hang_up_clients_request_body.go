package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HangUpClientsRequestBody struct {

	// 节点id
	NodeId string `json:"node_id"`

	// 要kill的会话addr列表
	ClientAddrs []string `json:"client_addrs"`
}

func (o HangUpClientsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpClientsRequestBody struct{}"
	}

	return strings.Join([]string{"HangUpClientsRequestBody", string(data)}, " ")
}
