package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiProcessNetInfo **参数解释** AI应用进程监听的网络信息
type AiProcessNetInfo struct {

	// **参数解释**： 应用进程监听IP **取值范围**： 取值0-2147483647
	ListenIp *string `json:"listen_ip,omitempty"`

	// **参数解释**： 应用进程监听对应的网络协议 **取值范围**： - tcp：tcp协议 - udp：udp协议
	ListenProtocol *string `json:"listen_protocol,omitempty"`

	// **参数解释**： 应用进程监听端口 **取值范围**： 取值0-2147483647
	ListenPort *int64 `json:"listen_port,omitempty"`

	// **参数解释**： 应用进程监听状态 **取值范围**： - established：已建立连接 - closed：连接已关闭 - listening：监听中 - other：连接中间态
	ListenStatus *string `json:"listen_status,omitempty"`
}

func (o AiProcessNetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiProcessNetInfo struct{}"
	}

	return strings.Join([]string{"AiProcessNetInfo", string(data)}, " ")
}
