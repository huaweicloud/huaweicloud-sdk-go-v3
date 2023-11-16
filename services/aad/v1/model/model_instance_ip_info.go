package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceIpInfo struct {

	// IP ID
	IpId *string `json:"ip_id,omitempty"`

	// IP
	Ip *string `json:"ip,omitempty"`

	// 保底带宽
	BasicBandwidth *int32 `json:"basic_bandwidth,omitempty"`

	// 弹性带宽
	ElasticBandwidth *int32 `json:"elastic_bandwidth,omitempty"`

	// IP状态
	IpStatus *int32 `json:"ip_status,omitempty"`
}

func (o InstanceIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceIpInfo struct{}"
	}

	return strings.Join([]string{"InstanceIpInfo", string(data)}, " ")
}
