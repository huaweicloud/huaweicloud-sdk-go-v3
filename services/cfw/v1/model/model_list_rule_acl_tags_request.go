package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleAclTagsRequest Request Object
type ListRuleAclTagsRequest struct {

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，配置后可根据企业项目过滤不同企业项目下的资产，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **约束限制**： 必须为数字 **取值范围**： 大于或等于0 **默认取值**： 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**： 每页显示个数 **约束限制**： 必须为数字 **取值范围**： 1-1024 **默认取值**： 不涉及
	Limit int32 `json:"limit"`
}

func (o ListRuleAclTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleAclTagsRequest struct{}"
	}

	return strings.Join([]string{"ListRuleAclTagsRequest", string(data)}, " ")
}
