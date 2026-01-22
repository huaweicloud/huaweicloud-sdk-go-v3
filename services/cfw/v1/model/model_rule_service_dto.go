package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleServiceDto RuleServiceDto
type RuleServiceDto struct {

	// **参数解释**： 服务输入类型，用于明确规则的服务输入类型。 **约束限制**： 不涉及 **取值范围**： 0为手动输入类型，1为自动输入类型 **默认取值**： 不涉及
	Type int32 `json:"type"`

	// **参数解释**： 服务协议类型，用于明确规则引用服务协议类型。 **约束限制**： type为0（手动类型）时，此处不能为空。 **取值范围**： 协议类型：TCP为6，UDP为17，ICMP为1，ICMPv6为58，Any为-1 **默认取值**： 不涉及
	Protocol *int32 `json:"protocol,omitempty"`

	// **参数解释**： 协议列表，用于明确规则引用协议列表。 **约束限制**： type为0（手动类型）时，此处不能为空。 **取值范围**： 协议类型：TCP为6，UDP为17，ICMP为1，ICMPv6为58，Any为-1 **默认取值**： 不涉及
	Protocols *[]int32 `json:"protocols,omitempty"`

	// **参数解释**： 源端口，会话发起方的端口。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SourcePort *string `json:"source_port,omitempty"`

	// **参数解释**： 目的端口，会话接收方的端口。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DestPort *string `json:"dest_port,omitempty"`

	// **参数解释**： 服务组ID，用于明确规则引用服务组，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 当address的type为1（关联IP地址组）时，此处不能为空 **取值范围**： 不涉及 **默认取值**： 不涉及
	ServiceSetId *string `json:"service_set_id,omitempty"`

	// **参数解释**： 服务（协议、源端口、目的端口）组的名称，用于明确规则引用服务组，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.name（.表示各对象之间层级的区分）获得。 **约束限制**： 当address的type为1（关联IP地址组）时，此处不能为空 **取值范围**： 不涉及 **默认取值**： 不涉及
	ServiceSetName *string `json:"service_set_name,omitempty"`

	// **参数解释**： 自定义服务，用于明确规则引用服务。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CustomService *[]ServiceItem `json:"custom_service,omitempty"`

	// **参数解释**： 预定义服务组ID列表，用于明确规则引用预定义服务组。服务组ID可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 查询条件中query_service_set_type需要设置为1预定义服务组。 **取值范围**： 不涉及 **默认取值**： 不涉及
	PredefinedGroup *[]string `json:"predefined_group,omitempty"`

	// **参数解释**： 服务组ID列表，用于明确规则引用服务组。服务组ID可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 查询条件中query_service_set_type需要设置为0自定义服务组。 **取值范围**： 不涉及 **默认取值**： 不涉及
	ServiceGroup *[]string `json:"service_group,omitempty"`

	// **参数解释**： 服务（协议、源端口、目的端口）组的名称。列表。 **约束限制**： 不涉及
	ServiceGroupNames *[]ServiceGroupVo `json:"service_group_names,omitempty"`

	// **参数解释**： 服务组类型，用于明确规则引用服务组类型。 **约束限制**： 不涉及 **取值范围**： 0表示自定义服务组，1表示常用Web服务，2表示常用远程登录和PING，3表示常用数据库 **默认取值**： 不涉及
	ServiceSetType *int32 `json:"service_set_type,omitempty"`
}

func (o RuleServiceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleServiceDto struct{}"
	}

	return strings.Join([]string{"RuleServiceDto", string(data)}, " ")
}
