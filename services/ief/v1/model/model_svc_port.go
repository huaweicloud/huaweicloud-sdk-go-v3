package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SvcPort 服务需要暴露的端口列表
type SvcPort struct {

	// 服务端口必须进行命名，而且名称只允许是{protocol}-{suffix}这种格式，其中{protocol}可以是tcp、http等，IEF根据在端口上定义的协议来提供对应的路由能力。例如“name:http-0”和“name:tcp-0”是合法的端口名，“name:http2forecast”是非法的端口号。
	Name *string `json:"name,omitempty"`

	// 当spec.type=NodePort时，指定映射到物理机的端口号
	NodePort *string `json:"node_port,omitempty"`

	// 服务监听的端口号
	Port *string `json:"port,omitempty"`

	// 具体的协议，比如TCP
	Protocol *string `json:"protocol,omitempty"`

	// 需要转发到后端Pod的端口号
	TargetPort *string `json:"target_port,omitempty"`
}

func (o SvcPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SvcPort struct{}"
	}

	return strings.Join([]string{"SvcPort", string(data)}, " ")
}
