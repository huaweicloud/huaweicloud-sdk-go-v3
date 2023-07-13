package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonPoolsWithBorderGroupDict 分组详情
type CommonPoolsWithBorderGroupDict struct {

	// 同组的公共池列表
	PublicipPools *[]string `json:"publicip_pools,omitempty"`

	// 分组：中心还是边缘
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o CommonPoolsWithBorderGroupDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonPoolsWithBorderGroupDict struct{}"
	}

	return strings.Join([]string{"CommonPoolsWithBorderGroupDict", string(data)}, " ")
}
