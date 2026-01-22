package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpRegionDto struct {

	// **参数解释**： 区域ID，用于明确规则引用地域，可通过[获取账号、IAM用户、项目、用户组、区域、委托的名称和ID](cfw_02_0030.xml)获取。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释**： 区域中文描述，仅当区域为中国区域时使用，可通过[地域信息表](cfw_02_0031.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DescriptionCn *string `json:"description_cn,omitempty"`

	// **参数解释**： 区域英文描述，仅当区域为非中国区域时使用，可通过[地域信息表](cfw_02_0031.xml)获取。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DescriptionEn *string `json:"description_en,omitempty"`

	// **参数解释**： 区域类型 **约束限制**： 不涉及 **取值范围**： - 0：国家 - 1：省份 - 2：大洲 **默认取值**： 不涉及
	RegionType *int32 `json:"region_type,omitempty"`
}

func (o IpRegionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpRegionDto struct{}"
	}

	return strings.Join([]string{"IpRegionDto", string(data)}, " ")
}
