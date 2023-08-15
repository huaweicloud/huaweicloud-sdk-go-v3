package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Region struct {

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 区域名称
	RegionName *string `json:"region_name,omitempty"`
}

func (o Region) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Region struct{}"
	}

	return strings.Join([]string{"Region", string(data)}, " ")
}
