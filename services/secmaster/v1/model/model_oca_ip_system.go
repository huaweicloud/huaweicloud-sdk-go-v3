package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIpSystem 系统信息
type OcaIpSystem struct {

	// 系统类型
	Family *string `json:"family,omitempty"`

	// 系统名称
	Name *string `json:"name,omitempty"`

	// 系统版本
	Version *string `json:"version,omitempty"`

	Vendor *OcaIpVendor `json:"vendor,omitempty"`
}

func (o OcaIpSystem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIpSystem struct{}"
	}

	return strings.Join([]string{"OcaIpSystem", string(data)}, " ")
}
