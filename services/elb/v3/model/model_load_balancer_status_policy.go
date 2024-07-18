package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatusPolicy LB状态树的转发策略状态信息
type LoadBalancerStatusPolicy struct {

	// 匹配后动作。  取值： - REDIRECT_TO_POOL：转发到后端服务器组。 - REDIRECT_TO_LISTENER：转发到监听器。
	Action *string `json:"action,omitempty"`

	// 转发策略ID。
	Id *string `json:"id,omitempty"`

	// 转发策略的配置状态。  取值范围： - ACTIVE: 默认值，表示正常。 [- ERROR: 表示当前策略与同一监听器下的其他策略存在相同的规则配置。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs)
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	// 转发策略名称。
	Name *string `json:"name,omitempty"`

	// 转发规则状态信息。
	Rules *[]LoadBalancerStatusL7Rule `json:"rules,omitempty"`
}

func (o LoadBalancerStatusPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPolicy struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPolicy", string(data)}, " ")
}
