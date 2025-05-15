package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionParam region信息。
type RegionParam struct {

	// 区域id。
	RegionId *string `json:"region_id,omitempty"`

	// 区域名称。
	RegionName *string `json:"region_name,omitempty"`

	// 区域。
	Region *string `json:"region,omitempty"`

	// 企业项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 区域状态。
	Status *string `json:"status,omitempty"`
}

func (o RegionParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionParam struct{}"
	}

	return strings.Join([]string{"RegionParam", string(data)}, " ")
}
