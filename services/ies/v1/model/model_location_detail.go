package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 场地详情
type LocationDetail struct {

	// 场地名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 场地描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 场地所在国家
	Country *string `json:"country,omitempty" xml:"country"`

	// 场地所在省/自治区/直辖市
	Province *string `json:"province,omitempty" xml:"province"`

	// 场地所在市/区
	City *string `json:"city,omitempty" xml:"city"`

	// 场地所在区/县
	District *string `json:"district,omitempty" xml:"district"`

	Condition *Condition `json:"condition,omitempty" xml:"condition"`
}

func (o LocationDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocationDetail struct{}"
	}

	return strings.Join([]string{"LocationDetail", string(data)}, " ")
}
