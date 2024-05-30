package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AllModelStatisticVo struct {

	// 是否为常用。
	Frequent *[]ModelStatisticVo `json:"frequent,omitempty"`

	// 首层模型。
	Top *[]ModelStatisticVo `json:"top,omitempty"`

	// 逻辑模型。
	Logic *[]ModelStatisticVo `json:"logic,omitempty"`

	// 物理模型。
	Physical *[]ModelStatisticVo `json:"physical,omitempty"`

	Dwr *ModelStatisticVo `json:"dwr,omitempty"`

	Dm *ModelStatisticVo `json:"dm,omitempty"`
}

func (o AllModelStatisticVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllModelStatisticVo struct{}"
	}

	return strings.Join([]string{"AllModelStatisticVo", string(data)}, " ")
}
