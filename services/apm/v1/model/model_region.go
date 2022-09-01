package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// region信息
type Region struct {

	// 区域id
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 区域名称
	RegionName *string `json:"region_name,omitempty" xml:"region_name"`

	// 区域
	Region *string `json:"region,omitempty" xml:"region"`

	// 企业项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 区域状态
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o Region) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Region struct{}"
	}

	return strings.Join([]string{"Region", string(data)}, " ")
}
