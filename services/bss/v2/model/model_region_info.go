package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionInfo 云服务区信息
type RegionInfo struct {

	// 云服务区编码
	RegionCode *string `json:"region_code,omitempty"`

	// 云服务区名称
	RegionName *string `json:"region_name,omitempty"`
}

func (o RegionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionInfo struct{}"
	}

	return strings.Join([]string{"RegionInfo", string(data)}, " ")
}
