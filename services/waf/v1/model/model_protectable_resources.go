package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectableResources struct {

	// **参数解释：** 负载均衡器名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LoadbalancerName *string `json:"loadbalancer_name,omitempty"`

	// **参数解释：** 负载均衡器ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// **参数解释：** 负载均衡器所属的账号ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 负载均衡器所在的项目ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 当前ELB所关联的监听器列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Listeners *[]Listener `json:"listeners,omitempty"`

	// **参数解释：** 负载均衡器绑定的EIP **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Eips *[]EipInfo `json:"eips,omitempty"`
}

func (o ProtectableResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableResources struct{}"
	}

	return strings.Join([]string{"ProtectableResources", string(data)}, " ")
}
