package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertEnvironment 告警产生的环境坐标信息
type AlertEnvironment struct {

	// 环境供应商
	VendorType *string `json:"vendor_type,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 区域id，全局服务global
	RegionId *string `json:"region_id,omitempty"`

	// 项目id， 全局服务默认null
	ProjectId *string `json:"project_id,omitempty"`
}

func (o AlertEnvironment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertEnvironment struct{}"
	}

	return strings.Join([]string{"AlertEnvironment", string(data)}, " ")
}
