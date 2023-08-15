package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateListenerQuicConfigOption 当前监听器关联的QUIC监听器配置信息，仅protocol为HTTPS时有效。  对于TCP/UDP/HTTP/QUIC监听器，若该字段非空则报错。  > 客户端向服务端发送正常的HTTP协议请求并携带了支持QUIC协议的信息。 如果服务端监听器开启了升级QUIC，那么就会在响应头中加入服务端支持的QUIC端口和版本信息。 客户端再次请求时会同时发送TCP(HTTPS)和UDP(QUIC)请求，若QUIC请求成功，则后续继续使用QUIC交互。  [不支持QUIC。](tag:hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
type CreateListenerQuicConfigOption struct {

	// 监听器关联的QUIC监听器ID。指定的listener id必须已存在，且协议类型为QUIC，不能指定为null，否则与enable_quic_upgrade冲突。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	QuicListenerId string `json:"quic_listener_id"`

	// QUIC升级的开启状态。 True:开启QUIC升级； Flase：关闭QUIC升级（默认）。 开启HTTPS监听器升级QUIC监听器能力。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	EnableQuicUpgrade *bool `json:"enable_quic_upgrade,omitempty"`
}

func (o CreateListenerQuicConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerQuicConfigOption struct{}"
	}

	return strings.Join([]string{"CreateListenerQuicConfigOption", string(data)}, " ")
}
