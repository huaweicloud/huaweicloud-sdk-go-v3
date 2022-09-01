package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LB状态树的转发策略状态信息
type LoadBalancerStatusPolicy struct {

	// 匹配后动作。取值： - REDIRECT_TO_POOL：转发到后端服务器组。 - REDIRECT_TO_LISTENER：转发到监听器。
	Action *string `json:"action,omitempty" xml:"action"`

	// 转发策略ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 转发策略的配置状态。取值： - ACTIVE：使用中，默认值。 - ERROR：表示当前策略与同一监听器下的其他策略存在相同的规则配置。
	ProvisioningStatus *string `json:"provisioning_status,omitempty" xml:"provisioning_status"`

	// 转发策略名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 转发规则状态信息。
	Rules *[]LoadBalancerStatusL7Rule `json:"rules,omitempty" xml:"rules"`
}

func (o LoadBalancerStatusPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPolicy struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPolicy", string(data)}, " ")
}
