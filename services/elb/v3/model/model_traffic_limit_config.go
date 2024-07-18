package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrafficLimitConfig 转发策略限速的配置。
type TrafficLimitConfig struct {

	// 转发策略整体限速。取值： 0-100000s。0表示不限速
	Qps *int32 `json:"qps,omitempty"`

	// 对转发策略单源进行限速。 quic监听器下转发策略不支持配置单源限速，指定该字段时，赋值可以为0或者为null 取值： 0-100000s。0表示不限速，如果qps不为0，per_ip_qps需要小于qps。
	PerSourceIpQps *int32 `json:"per_source_ip_qps,omitempty"`

	// 对转发策略单源进行限速。取值： 0-100000s。当qps超限的时候，不返回503，支持允许局部突增burst大小的请求。
	Burst *int32 `json:"burst,omitempty"`
}

func (o TrafficLimitConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficLimitConfig struct{}"
	}

	return strings.Join([]string{"TrafficLimitConfig", string(data)}, " ")
}
