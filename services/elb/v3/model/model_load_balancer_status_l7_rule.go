package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询负载均衡状态树返回对象中的rule数据模型。
type LoadBalancerStatusL7Rule struct {

	// L7转发规则ID。
	Id string `json:"id"`

	// 匹配内容类型。取值： - HOST_NAME：域名匹配。 - PATH：URL路径匹配。  使用说明： 同一个l7policy下创建的所有的l7rule的HOST_NAME不能重复。
	Type string `json:"type"`

	// 转发规则的配置状态。 取值： - ACTIVE：使用中，默认值。 - ERROR：当前规则所属策略与同一监听器下的其他策略存在相同的规则配置。
	ProvisioningStatus string `json:"provisioning_status"`
}

func (o LoadBalancerStatusL7Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusL7Rule struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusL7Rule", string(data)}, " ")
}
