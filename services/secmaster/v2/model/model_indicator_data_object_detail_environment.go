package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorDataObjectDetailEnvironment 环境信息
type IndicatorDataObjectDetailEnvironment struct {

	// 环境供应商（如HWC,AWS,Azure等）
	VendorType *string `json:"vendor_type,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`
}

func (o IndicatorDataObjectDetailEnvironment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorDataObjectDetailEnvironment struct{}"
	}

	return strings.Join([]string{"IndicatorDataObjectDetailEnvironment", string(data)}, " ")
}
