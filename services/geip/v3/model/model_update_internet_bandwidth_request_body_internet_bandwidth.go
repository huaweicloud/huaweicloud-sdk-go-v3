package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInternetBandwidthRequestBodyInternetBandwidth 更新全域公网带宽请求对象
type UpdateInternetBandwidthRequestBodyInternetBandwidth struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
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
