package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachInternalIpRequestBody struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	// 新的Ip需要为用户可用vpc中的网段。只支持IPV4。
	NewIp string `json:"new_ip"`
}

func (o AttachInternalIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternalIpRequestBody struct{}"
	}

	return strings.Join([]string{"AttachInternalIpRequestBody", string(data)}, " ")
}
