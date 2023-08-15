package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocationDetail 场地详情
type LocationDetail struct {

	// 场地名称（已废弃），传入该参数不会再生效，新建站点也不会再返回该字段
	Name *string `json:"name,omitempty"`

	// 场地描述
	Description *string `json:"description,omitempty"`

	// 场地所在国家
	Country *string `json:"country,omitempty"`

	// 场地所在省/自治区/直辖市
	Province *string `json:"province,omitempty"`

	// 场地所在市/区
	City *string `json:"city,omitempty"`

	// 场地所在区/县
	District *string `json:"district,omitempty"`

	Condition *Condition `json:"condition,omitempty"`
}

func (o LocationDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocationDetail struct{}"
	}

	return strings.Join([]string{"LocationDetail", string(data)}, " ")
}
