package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DmsKafkaInfo struct {

	// Kafka连接地址
	Addr *[]string `json:"addr,omitempty" xml:"addr"`

	// Kafka模式下，是否开启共识节点老化
	OrderFadeEnable *bool `json:"order_fade_enable,omitempty" xml:"order_fade_enable"`

	// Kafka模式下，开启共识节点后的老化阈值
	OrderFadeCache *int64 `json:"order_fade_cache,omitempty" xml:"order_fade_cache"`
}

func (o DmsKafkaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DmsKafkaInfo struct{}"
	}

	return strings.Join([]string{"DmsKafkaInfo", string(data)}, " ")
}
