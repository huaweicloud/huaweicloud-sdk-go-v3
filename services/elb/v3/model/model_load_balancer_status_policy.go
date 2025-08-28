package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatusPolicy **参数解释**：监听器下的7层转发策略的状态信息。
type LoadBalancerStatusPolicy struct {

	// **参数解释**：匹配后动作。  **取值范围**： - REDIRECT_TO_POOL：转发到后端服务器组。 - REDIRECT_TO_LISTENER：转发到监听器。
	Action *string `json:"action,omitempty"`

	// **参数解释**：转发策略ID。  **取值范围**：不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：转发策略的配置状态。  **取值范围**： - ACTIVE: 默认值，表示正常。 [- ERROR: 表示当前策略与同一监听器下的其他策略存在相同的规则配置。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs)
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	// **参数解释**：转发策略名称。  **取值范围**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：当前转发策略下的所有转发规则状的态信息。
	Rules *[]LoadBalancerStatusL7Rule `json:"rules,omitempty"`
}

func (o LoadBalancerStatusPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPolicy struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPolicy", string(data)}, " ")
}
