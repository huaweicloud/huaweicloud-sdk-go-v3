package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIpService 应用信息
type OcaIpService struct {

	// 应用对应端口
	Port *int32 `json:"port,omitempty"`

	// 协议
	Protocol *string `json:"protocol,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用版本
	Version *string `json:"version,omitempty"`

	Vendor *OcaIpVendor `json:"vendor,omitempty"`
}

func (o OcaIpService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIpService struct{}"
	}

	return strings.Join([]string{"OcaIpService", string(data)}, " ")
}
