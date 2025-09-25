package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetWtpProtectionStatusRequestInfo struct {

	// **参数解释**: 防护状态 **约束限制**: 不涉及 **取值范围**: - True ：开启防护，必须填写charging_mode。 - False ：关闭防护，无需填写charging_mode。  **默认取值**: False
	Status bool `json:"status"`

	// **参数解释**: 需要开启或关闭防护的服务器ID列表。 **约束限制** : 开启防护时，需要使用 ListWtpProtectHost 接口查询网页防篡改主机防护状态列表信息，在 ListWtpProtectHost 接口的响应体中，protect_status 等于 closed 且 agent_status 等于 online 的 host_id 是符合开启防护的服务器ID。 **取值范围**: 最少1条，最多20000条 **默认取值** : 不涉及
	HostIdList []string `json:"host_id_list"`

	// **参数解释**: 计费模式 **约束限制**: 不涉及 **取值范围**: - packet_cycle: 包年/包月，可填写resource_id。 - on_demand: 按需计费，无需填写resource_id。  **默认取值**: on_demand
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**: 资源ID，即网页防篡改配额的配额ID，当charging_mode选择packet_cycle时可填写该字段，表示使用一个指定配额，也可不填写该字段，表示随机选择符合的配额。 **约束限制** : 不涉及 **取值范围**: 字符长度0-64位 **默认取值** : 不涉及
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o SetWtpProtectionStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetWtpProtectionStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SetWtpProtectionStatusRequestInfo", string(data)}, " ")
}
