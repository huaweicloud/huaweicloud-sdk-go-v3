package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpRegionDto struct {

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 中文描述
	DescriptionCn *string `json:"description_cn,omitempty"`

	// 英文描述
	DescriptionEn *string `json:"description_en,omitempty"`

	// 区域类型，0表示国家，1表示省份，2表示大洲
	RegionType *int32 `json:"region_type,omitempty"`
}

func (o IpRegionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpRegionDto struct{}"
	}

	return strings.Join([]string{"IpRegionDto", string(data)}, " ")
}
