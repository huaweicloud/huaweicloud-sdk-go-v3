package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainSetsRequest Request Object
type ListDomainSetsRequest struct {

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，配置后可根据企业项目过滤不同企业项目下的资产，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 每页显示个数 **约束限制**： 不涉及 **取值范围**： 1-1024 **默认取值**： 不涉及
	Limit int32 `json:"limit"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **约束限制**： 不涉及 **取值范围**： 大于等于0 **默认取值**： 0
	Offset int32 `json:"offset"`

	// **参数解释**： 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得 **约束限制**： type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID。此处仅取type为1的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。 **取值范围**： 不涉及 **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： 关键字，可使用域名组名称或描述 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	KeyWord *string `json:"key_word,omitempty"`

	// **参数解释**： 域名组类型 **约束限制**： 不涉及 **取值范围**： - 0：应用域名组 - 1：网络域名组 **默认取值**： 不涉及
	DomainSetType *int32 `json:"domain_set_type,omitempty"`

	// **参数解释**： 配置状态 **约束限制**： 不涉及 **取值范围**： - -1：未配置态 - 0：配置失败 - 1：配置成功 - 2：配置中 - 3：正常 - 4：配置异常 **默认取值**： 不涉及
	ConfigStatus *int32 `json:"config_status,omitempty"`
}

func (o ListDomainSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainSetsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainSetsRequest", string(data)}, " ")
}
