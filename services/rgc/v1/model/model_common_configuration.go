package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonConfiguration Landing Zone配置信息。
type CommonConfiguration struct {

	// 主区域名。
	HomeRegion *string `json:"home_region,omitempty"`

	// CTS配置状态。
	CloudTrailType *bool `json:"cloud_trail_type,omitempty"`

	// 是否设置Landing Zone的identity center。
	IdentityCenterStatus *string `json:"identity_center_status,omitempty"`

	// 设置organization的类型。STANDARD和NON_STANDARD。
	OrganizationStructureType *string `json:"organization_structure_type,omitempty"`
}

func (o CommonConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonConfiguration struct{}"
	}

	return strings.Join([]string{"CommonConfiguration", string(data)}, " ")
}
