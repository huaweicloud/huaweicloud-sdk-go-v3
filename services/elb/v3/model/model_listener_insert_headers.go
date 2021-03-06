package model

import (
	"encoding/json"

	"strings"
)

// 可选HTTP头插入。
type ListenerInsertHeaders struct {
	// X-Forwarded-ELB-IP设为true可以将ELB实例的eip地址从报文的http头中带到后端云服务器。

	XForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`
	// X-Forwarded-Port设为true可以将ELB实例的监听端口从报文的http头中带到后端云服务器。

	XForwardedPort *bool `json:"X-Forwarded-Port,omitempty"`
	// X-Forwarded-For-Port设为true可以将客户端的源端口从报文的http头中带到后端云服务器。

	XForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`
	// X-Forwarded-Host设为true可以将客户请求头的X-Forwarded-Host设置为请求头的Host带到后端云服务器。

	XForwardedHost bool `json:"X-Forwarded-Host"`
}

func (o ListenerInsertHeaders) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListenerInsertHeaders struct{}"
	}

	return strings.Join([]string{"ListenerInsertHeaders", string(data)}, " ")
}
