package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchFirewallEipProtectionRespData 一键逃生/一键恢复返回数据
type SwitchFirewallEipProtectionRespData struct {

	// **参数解释**： 防火墙防护状态，0: 正常状态, 1: bypass进行中, 2: bypass成功, 3: bypass失败, 4: 恢复中, 5: 恢复失败 **取值范围**： 不涉及
	ProtectionStatus *int32 `json:"protection_status,omitempty"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId *string `json:"object_id,omitempty"`
}

func (o SwitchFirewallEipProtectionRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFirewallEipProtectionRespData struct{}"
	}

	return strings.Join([]string{"SwitchFirewallEipProtectionRespData", string(data)}, " ")
}
