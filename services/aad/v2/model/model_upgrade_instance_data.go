package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeInstanceData struct {

	// 保底带宽(G)
	BasicBandwidth *string `json:"basic_bandwidth,omitempty"`

	// 弹性带宽(G)
	ElasticBandwidth *string `json:"elastic_bandwidth,omitempty"`

	// 业务带宽
	ServiceBandwidth *int32 `json:"service_bandwidth,omitempty"`

	// 端口数
	PortNum *int32 `json:"port_num,omitempty"`

	// 域名数
	BindDomainNum *int32 `json:"bind_domain_num,omitempty"`

	// 弹性业务带宽,0-关闭，3-月95
	ElasticServiceBandwidthType *int32 `json:"elastic_service_bandwidth_type,omitempty"`

	// 弹性业务带宽增加值
	ElasticServiceBandwidth *int32 `json:"elastic_service_bandwidth,omitempty"`
}

func (o UpgradeInstanceData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceData struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceData", string(data)}, " ")
}
