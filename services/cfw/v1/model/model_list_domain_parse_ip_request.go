package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseIpRequest Request Object
type ListDomainParseIpRequest struct {

	// **参数解释**： 地址类型 **约束限制**： 不涉及 **取值范围**： - 0：ipv4 - 1：ipv6 **默认取值**： 不涉及
	AddressType *int32 `json:"address_type,omitempty"`

	// **参数解释**： 域名id，域名id可通过[获取域名组下域名列表接口](ListDomains.xml)查询获得，通过返回值中的data.records.domain_address_id（.表示各对象之间层级的区分）获取 **约束限制**： 仅网络型域名组需填写此参数 **取值范围**： 不涉及 **默认取值**： 不涉及
	DomainAddressId string `json:"domain_address_id"`

	// **参数解释**： 域名组ID，可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获取 **约束限制**： 仅网络型域名组需填写此参数 **取值范围**： 不涉及 **默认取值**： 不涉及
	DomainSetId string `json:"domain_set_id"`

	// **参数解释**： 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，配置后可根据企业项目过滤不同企业项目下的资产，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListDomainParseIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainParseIpRequest struct{}"
	}

	return strings.Join([]string{"ListDomainParseIpRequest", string(data)}, " ")
}
