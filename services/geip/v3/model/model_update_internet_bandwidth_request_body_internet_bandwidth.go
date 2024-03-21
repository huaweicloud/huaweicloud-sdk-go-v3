package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInternetBandwidthRequestBodyInternetBandwidth 更新全域公网带宽请求对象
type UpdateInternetBandwidthRequestBodyInternetBandwidth struct {

	// - 功能说明：全域公网带宽名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// - 功能说明：用户自定义的资源描述 - 约束：   - 值的长度最大512字符，由数字、字母、中文、_(下划线)、-（中划线）、.（点）组成。
	Description *string `json:"description,omitempty"`

	// 全域公网带宽大小（出云方向）
	Size *int32 `json:"size,omitempty"`

	// 计费模式
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 全域公网带宽大小（入云方向）
	IngressSize *int32 `json:"ingress_size,omitempty"`
}

func (o UpdateInternetBandwidthRequestBodyInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternetBandwidthRequestBodyInternetBandwidth struct{}"
	}

	return strings.Join([]string{"UpdateInternetBandwidthRequestBodyInternetBandwidth", string(data)}, " ")
}
