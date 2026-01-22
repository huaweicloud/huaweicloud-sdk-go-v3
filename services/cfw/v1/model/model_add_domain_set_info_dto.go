package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDomainSetInfoDto struct {

	// **参数解释**： 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得 **约束限制**： type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID。此处仅取type为1的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。 **取值范围**： 不涉及 **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： 域名组名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 域名组描述 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 域名信息列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DomainNames []DomainSetInfoDto `json:"domain_names"`

	// **参数解释**： 域名组类型 **约束限制**： 不涉及 **取值范围**： - 0：应用域名组 - 1：网络域名组 **默认取值**： 不涉及
	DomainSetType *int32 `json:"domain_set_type,omitempty"`
}

func (o AddDomainSetInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainSetInfoDto struct{}"
	}

	return strings.Join([]string{"AddDomainSetInfoDto", string(data)}, " ")
}
