package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicIpBandwidthOption 创建弹性公网IP的带宽参数
type CreatePublicIpBandwidthOption struct {

	// 带宽id。
	Id *string `json:"id,omitempty"`

	// 带宽类型。  指定带宽ID，则该参数不生效。 不指定带宽的情况下，如果当前带宽类型下没有带宽，自动在该带宽类型下创建带宽，有则使用最近创建的带宽。
	BandwidthType *string `json:"bandwidth_type,omitempty"`
}

func (o CreatePublicIpBandwidthOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicIpBandwidthOption struct{}"
	}

	return strings.Join([]string{"CreatePublicIpBandwidthOption", string(data)}, " ")
}
