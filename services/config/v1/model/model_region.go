package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Region 区域
type Region struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 区域显示名称
	DisplayName *string `json:"display_name,omitempty"`
}

func (o Region) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Region struct{}"
	}

	return strings.Join([]string{"Region", string(data)}, " ")
}
