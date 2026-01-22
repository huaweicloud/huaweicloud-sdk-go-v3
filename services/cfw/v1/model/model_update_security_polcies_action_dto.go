package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityPolciesActionDto struct {

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志ID，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得 **约束限制**： type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： 规则动作 **约束限制**： 不涉及 **取值范围**： enable表示允许通行（permit），disable表示拒绝通行（deny） **默认取值**： 不涉及
	Action string `json:"action"`

	// **参数解释**： 规则ID列表，规则ID可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RuleIds []string `json:"rule_ids"`
}

func (o UpdateSecurityPolciesActionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolciesActionDto struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolciesActionDto", string(data)}, " ")
}
