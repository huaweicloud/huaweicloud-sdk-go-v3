package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Ipv6Bandwidth struct {

	// IPv6带宽的ID。
	Id *string `json:"id,omitempty"`

	// 带宽类型。  指定带宽ID，则该参数不生效。 不指定带宽的情况下，如果当前带宽类型下没有带宽，自动在该带宽类型下创建带宽，有则使用最近创建的带宽。 约束：指定的bandwidth_type必须在对应弹性公网IP类型的allow_share_bandwidth_types中才能使用
	BandwidthType *string `json:"bandwidth_type,omitempty"`
}

func (o Ipv6Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ipv6Bandwidth struct{}"
	}

	return strings.Join([]string{"Ipv6Bandwidth", string(data)}, " ")
}
