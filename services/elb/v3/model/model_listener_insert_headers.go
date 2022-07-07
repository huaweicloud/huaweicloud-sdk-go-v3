package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可选的HTTP头插入，可以将从负载均衡器到后端云服务器的路径中需要被后端云服务器用到的信息写入HTTP中，随报文传递到后端云服务器使。例如可通过X-Forwarded-ELB-IP开关，将负载均衡器的弹性公网IP传到后端云服务器。
type ListenerInsertHeaders struct {

	// X-Forwarded-ELB-IP设为true可以将ELB实例的eip地址从报文的http头中带到后端云服务器。
	XForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`

	// X-Forwarded-Port设为true可以将ELB实例的监听端口从报文的http头中带到后端云服务器。
	XForwardedPort *bool `json:"X-Forwarded-Port,omitempty"`

	// X-Forwarded-For-Port设为true可以将客户端的源端口从报文的http头中带到后端云服务器。
	XForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`

	// X-Forwarded-Host设为true可以将客户请求头的X-Forwarded-Host设置为请求头的Host带到后端云服务器。
	XForwardedHost *bool `json:"X-Forwarded-Host,omitempty"`
}

func (o ListenerInsertHeaders) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerInsertHeaders struct{}"
	}

	return strings.Join([]string{"ListenerInsertHeaders", string(data)}, " ")
}
