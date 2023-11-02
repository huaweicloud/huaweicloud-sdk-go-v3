package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorDetailEnvironment 环境信息
type CreateIndicatorDetailEnvironment struct {

	// 环境供应商，如：HWC/AWS等
	VendorType string `json:"vendor_type"`

	// 租户ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 项目ID
	ProjectId string `json:"project_id"`
}

func (o CreateIndicatorDetailEnvironment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorDetailEnvironment struct{}"
	}

	return strings.Join([]string{"CreateIndicatorDetailEnvironment", string(data)}, " ")
}
