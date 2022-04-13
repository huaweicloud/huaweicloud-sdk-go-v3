package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Region struct {
	// 区域描述信息。

	Description string `json:"description"`
	// null.

	ParentRegionId string `json:"parent_region_id"`

	Links *LinksSelf `json:"links"`

	Locales *RegionLocales `json:"locales"`
	// 区域ID。

	Id string `json:"id"`
	// 区域类型。

	Type string `json:"type"`
}

func (o Region) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Region struct{}"
	}

	return strings.Join([]string{"Region", string(data)}, " ")
}
