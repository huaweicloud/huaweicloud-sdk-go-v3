package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseAppBasicConfig 应用基础配置。
type SupabaseAppBasicConfig struct {

	// 区域ID。
	Region string `json:"region"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 子网网络ID。
	SubnetId string `json:"subnet_id"`

	// UTC时区（可选）。
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o SupabaseAppBasicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseAppBasicConfig struct{}"
	}

	return strings.Join([]string{"SupabaseAppBasicConfig", string(data)}, " ")
}
