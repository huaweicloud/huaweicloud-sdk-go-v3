package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryFireWallInstanceDto struct {

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 查询关键字，可为防火墙ID或防火墙名称的一部分。可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	KeyWord *string `json:"key_word,omitempty"`

	// **参数解释**： 标签列表，可通过查询标签服务查询标签接口获得，返回值即为标签列表 **约束限制**： 不涉及
	Tags *[]TagInfo `json:"tags,omitempty"`

	// **参数解释**： 每页显示个数 **约束限制**： 不涉及 **取值范围**： 1-1024 **默认取值**： 不涉及
	Limit int32 `json:"limit"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **约束限制**： 不涉及 **取值范围**： 大于或等于0 **默认取值**： 0
	Offset int32 `json:"offset"`
}

func (o QueryFireWallInstanceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryFireWallInstanceDto struct{}"
	}

	return strings.Join([]string{"QueryFireWallInstanceDto", string(data)}, " ")
}
