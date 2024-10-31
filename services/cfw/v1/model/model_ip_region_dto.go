package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpRegionDto struct {

	// 区域id，可通过[获取账号、IAM用户、项目、用户组、区域、委托的名称和ID](cfw_02_0030.xml)获取。
	RegionId *string `json:"region_id,omitempty"`

	// 区域中文描述，仅当区域为中国区域时使用，可通过[地域信息表](cfw_02_0031.xml)获取。
	DescriptionCn *string `json:"description_cn,omitempty"`

	// 区域英文描述，仅当区域为非中国区域时使用，可通过[地域信息表](cfw_02_0031.xml)获取。
	DescriptionEn *string `json:"description_en,omitempty"`

	// 区域类型，0表示国家，1表示省份，2表示大洲，可通过[地域信息表](cfw_02_0031.xml)获取。
	RegionType *int32 `json:"region_type,omitempty"`
}

func (o IpRegionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpRegionDto struct{}"
	}

	return strings.Join([]string{"IpRegionDto", string(data)}, " ")
}
