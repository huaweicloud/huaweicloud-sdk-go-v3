package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindingVifDetails
type BindingVifDetails struct {

	// 功能说明：取值为true，表示是虚拟机的主网卡。
	PrimaryInterface *bool `json:"primary_interface,omitempty"`

	// 功能说明：表示该网络服务提供端口过滤特性，如安全组和反MAC/IP欺骗。
	PortFilter *bool `json:"port_filter,omitempty"`

	// 用于通知像nova这样的API消费者，应该使用OVS的混合插入策略。
	OvsHybridPlug *bool `json:"ovs_hybrid_plug,omitempty"`
}

func (o BindingVifDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingVifDetails struct{}"
	}

	return strings.Join([]string{"BindingVifDetails", string(data)}, " ")
}
