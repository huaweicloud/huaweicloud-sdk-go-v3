package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFirewallReq 按需防火墙实体
type CreateFirewallReq struct {

	// **参数解释**： 防火墙名称 约束限制： 不涉及 **取值范围**： 不涉及 默认取值： 不涉及
	Name string `json:"name"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 约束限制： 不涉及 **取值范围**： 不涉及 默认取值： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 服务资源标签列表，防火墙资源添加标签后，可根据键、值组合查询资源，同时可根据键、值组合进行话单合并统计 **约束限制**： 不涉及
	Tags *[]CreateFirewallReqTags `json:"tags,omitempty"`

	Flavor *CreateFirewallReqFlavor `json:"flavor"`

	ChargeInfo *CreateFirewallReqChargeInfo `json:"charge_info"`
}

func (o CreateFirewallReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallReq struct{}"
	}

	return strings.Join([]string{"CreateFirewallReq", string(data)}, " ")
}
