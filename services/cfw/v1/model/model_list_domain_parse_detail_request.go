package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseDetailRequest Request Object
type ListDomainParseDetailRequest struct {

	// **参数解释**： 域名 **约束限制**： 域名格式，如www.example.com **取值范围**： 不涉及 **默认取值**： 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，配置后可根据企业项目过滤不同企业项目下的资产，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 参数解释： IP地址的互联网协议类型，用于指定IP地址的互联网协议，由客户指定 约束限制： 不涉及 取值范围： 0：IPv4 1：IPv6 默认取值： 不涉及
	AddressType *int32 `json:"address_type,omitempty"`
}

func (o ListDomainParseDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainParseDetailRequest struct{}"
	}

	return strings.Join([]string{"ListDomainParseDetailRequest", string(data)}, " ")
}
