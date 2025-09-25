package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Elbv3Listener struct {

	// 监听器ID。
	Id *string `json:"id,omitempty"`

	// 监听器的名称。
	Name *string `json:"name,omitempty"`

	// 监听器的监听协议。
	Protocol *string `json:"protocol,omitempty"`

	// 监听器的前端监听端口。
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	Ipgroup *ListenerIpGroup `json:"ipgroup,omitempty"`
}

func (o Elbv3Listener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Elbv3Listener struct{}"
	}

	return strings.Join([]string{"Elbv3Listener", string(data)}, " ")
}
