package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegionDetail struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 区域名称
	DisplayName *string `json:"display_name,omitempty"`
}

func (o RegionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionDetail struct{}"
	}

	return strings.Join([]string{"RegionDetail", string(data)}, " ")
}
