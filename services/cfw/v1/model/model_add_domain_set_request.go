package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDomainSetRequest Request Object
type AddDomainSetRequest struct {

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，字段废弃 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	Body *AddDomainSetInfoDto `json:"body,omitempty"`
}

func (o AddDomainSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainSetRequest struct{}"
	}

	return strings.Join([]string{"AddDomainSetRequest", string(data)}, " ")
}
