package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindingVifDetails
type BindingVifDetails struct {

	// 参数解释：表示该网卡是否为云服务器的主网卡。取值范围：true，主网卡；false，非主网卡。
	PrimaryInterface *bool `json:"primary_interface,omitempty"`

	// 参数解释：表示该网络服务提供端口过滤特性，如安全组和反MAC/IP欺骗。取值范围：true，提供端口过滤特性；false，未提供端口过滤特性。
	PortFilter *bool `json:"port_filter,omitempty"`

	// 参数解释：是否为ovs/bridge混合模式。取值范围：true，ovs/bridge混合模式；false，非ovs/bridge混合模式。
	OvsHybridPlug *bool `json:"ovs_hybrid_plug,omitempty"`
}

func (o BindingVifDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingVifDetails struct{}"
	}

	return strings.Join([]string{"BindingVifDetails", string(data)}, " ")
}
