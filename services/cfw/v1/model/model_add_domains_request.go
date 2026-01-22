package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDomainsRequest Request Object
type AddDomainsRequest struct {

	// **参数解释**： 域名组id，可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SetId string `json:"set_id"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，配置后可根据企业项目过滤不同企业项目下的资产，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	Body *AddDomainListDto `json:"body,omitempty"`
}

func (o AddDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainsRequest struct{}"
	}

	return strings.Join([]string{"AddDomainsRequest", string(data)}, " ")
}
