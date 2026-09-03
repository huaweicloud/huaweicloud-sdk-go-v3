package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtocolTestVo struct {
	Http *HttpVo `json:"http,omitempty"`

	// ping拨测任务信息
	Ping *[]PingVo `json:"ping,omitempty"`

	// ping/http节点地址
	PointHost *[]string `json:"point_host,omitempty"`

	// 协议
	Protocol *string `json:"protocol,omitempty"`
}

func (o ProtocolTestVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtocolTestVo struct{}"
	}

	return strings.Join([]string{"ProtocolTestVo", string(data)}, " ")
}
