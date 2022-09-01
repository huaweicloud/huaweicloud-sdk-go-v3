package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// listener对象中的quic配置信息，仅protocol为HTTPS时有效。 支持创建和修改； 支持HTTPS监听器升级QUIC监听器能力。仅HTTPS监听器支持升级到QUIC监听器 当客户开启升级之后选择关联的quic监听器，https对象要保存改quic监听器ID。 对于TCP/UDP/HTTP/QUIC监听器，若该字段非空则报错。
type UpdateListenerQuicConfigOption struct {

	// 监听器关联的QUIC监听器ID。指定的listener id必须已存在，且协议类型为QUIC，不能指定为null，否则与enable_quic_upgrade冲突。
	QuicListenerId *string `json:"quic_listener_id,omitempty" xml:"quic_listener_id"`

	// QUIC升级的开启状态。 True:开启QUIC升级； Flase：关闭QUIC升级； 开启HTTPS监听器升级QUIC监听器能力
	EnableQuicUpgrade *bool `json:"enable_quic_upgrade,omitempty" xml:"enable_quic_upgrade"`
}

func (o UpdateListenerQuicConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerQuicConfigOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerQuicConfigOption", string(data)}, " ")
}
