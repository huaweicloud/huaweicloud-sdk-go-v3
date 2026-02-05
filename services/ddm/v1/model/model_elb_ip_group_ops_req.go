package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ElbIpGroupOpsReq struct {

	// **参数解释**：  弹性负载均衡ip组名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  弹性负载均衡ip组类型。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：  是否开启弹性负载均衡ip组。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EnableIpGroup *bool `json:"enable_ip_group,omitempty"`

	// **参数解释**：  弹性负载均衡ip组id。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**：  ip列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	IpList *[]IpGroupItem `json:"ip_list,omitempty"`
}

func (o ElbIpGroupOpsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElbIpGroupOpsReq struct{}"
	}

	return strings.Join([]string{"ElbIpGroupOpsReq", string(data)}, " ")
}
