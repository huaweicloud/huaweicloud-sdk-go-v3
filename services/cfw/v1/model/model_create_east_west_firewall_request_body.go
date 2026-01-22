package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEastWestFirewallRequestBody **参数解释**： 创建东西向防火墙body体 **约束限制**： 不涉及
type CreateEastWestFirewallRequestBody struct {

	// **参数解释**： 出方向关联ER ID,可通过ER服务查询企业路由器列表接口获得，返回值中instances.id即为erid（.表示各对象之间层级的区分） **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ErId string `json:"er_id"`

	// **参数解释**： 创建引流VPC时使用的网段 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	InspectionCidr string `json:"inspection_cidr"`

	// **参数解释**： 东西向防火墙模式，填写er **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Mode string `json:"mode"`
}

func (o CreateEastWestFirewallRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEastWestFirewallRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEastWestFirewallRequestBody", string(data)}, " ")
}
