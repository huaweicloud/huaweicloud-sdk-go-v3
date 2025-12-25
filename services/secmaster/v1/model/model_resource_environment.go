package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceEnvironment 资产所属环境
type ResourceEnvironment struct {

	// 环境供应商
	VendorType string `json:"vendor_type"`

	// HWC special，资产归属
	DomainId string `json:"domain_id"`

	// HWC special，全局服务\"global\"，资产归属
	RegionId *string `json:"region_id,omitempty"`

	// HWC special， 全局服务默认null， 资产归属
	ProjectId *string `json:"project_id,omitempty"`

	// HWC special，资产归属的企业项目id。
	EpId *string `json:"ep_id,omitempty"`

	// HWC special，资产归属的企业项目名称。
	EpName *string `json:"ep_name,omitempty"`

	// OCA special，包含资产探针或资产提供商
	VendorName string `json:"vendor_name"`

	// OCA special，线下机房名称
	IdcName string `json:"idc_name"`

	// OCA special，线下机房id
	IdcId *string `json:"idc_id,omitempty"`
}

func (o ResourceEnvironment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceEnvironment struct{}"
	}

	return strings.Join([]string{"ResourceEnvironment", string(data)}, " ")
}
