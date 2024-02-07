package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowedBandwidthTypes 地址池支持的全域公网带宽类型资源
type AllowedBandwidthTypes struct {

	// 全域公网带宽类型名称
	Type *string `json:"type,omitempty"`

	// 中文名称
	CnName *string `json:"cn_name,omitempty"`

	// 英文名称
	EnName *string `json:"en_name,omitempty"`
}

func (o AllowedBandwidthTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowedBandwidthTypes struct{}"
	}

	return strings.Join([]string{"AllowedBandwidthTypes", string(data)}, " ")
}
