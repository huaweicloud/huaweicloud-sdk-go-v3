package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatusL7Rule **参数解释**：查询负载均衡状态树返回对象中的rule数据模型。  **默认取值**：不涉及
type LoadBalancerStatusL7Rule struct {

	// **参数解释**：L7转发规则ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：匹配内容类型。  **取值范围**： - HOST_NAME：域名匹配。 - PATH：URL路径匹配。
	Type string `json:"type"`

	// **参数解释**：转发规则的配置状态。  **取值范围**： - ACTIVE：使用中，默认值。 - ERROR：当前规则所属策略与同一监听器下的其他策略存在相同的规则配置。
	ProvisioningStatus string `json:"provisioning_status"`
}

func (o LoadBalancerStatusL7Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusL7Rule struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusL7Rule", string(data)}, " ")
}
