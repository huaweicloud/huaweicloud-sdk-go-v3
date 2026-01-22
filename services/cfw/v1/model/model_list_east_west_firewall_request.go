package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEastWestFirewallRequest Request Object
type ListEastWestFirewallRequest struct {

	// **参数解释**： 每页显示个数 **约束限制**： 不涉及 **取值范围**： 1-1024 **默认取值**： 不涉及
	Limit int32 `json:"limit"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **约束限制**： 必须为数字 **取值范围**： 大于或等于0 **默认取值**： 默认0
	Offset int32 `json:"offset"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 用户未开启企业项目时为0 **取值范围**： 不涉及 **默认取值**： 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`
}

func (o ListEastWestFirewallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEastWestFirewallRequest struct{}"
	}

	return strings.Join([]string{"ListEastWestFirewallRequest", string(data)}, " ")
}
