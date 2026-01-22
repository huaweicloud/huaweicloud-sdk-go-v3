package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FirewallInstanceResource struct {

	// **参数解释**： 资源id，包括防火墙资源id，带宽资源id，eip资源id，vpc资源id，cbc回调后返回id。 **取值范围**： 不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 服务类型，用于CBC使用，特指：hws.service.type.cfw **取值范围**： 不涉及
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// **参数解释**： 资源类型 **取值范围**： - 1：云防火墙:hws.resource.type.cfw - 2：EIP:hws.resource.type.cfw.exp.eip - 3：带宽:hws.resource.type.cfw.exp.bandwidth - 4：VPC:hws.resource.type.cfw.exp
	ResourceType *string `json:"resource_type,omitempty"`

	// **参数解释**： 库存单位码 **取值范围**： - 1：防火墙标准版cfw.standard - 2：防火墙专业版cfw.professional - 3：eip标准版cfw.expack.eip.standard - 4：eip专业版cfw.expack.eip.professional - 5：带宽基础版cfw.expack.bandwidth.standard - 6：带宽专业版cfw.expack.bandwidth.professional - 7：vpc专业版cfw.expack.vpc.professional
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// **参数解释**： 资源数量 **取值范围**： 不涉及
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// **参数解释**： 资源单位 **取值范围**： 不涉及
	ResourceSizeMeasureId *int32 `json:"resource_size_measure_id,omitempty"`
}

func (o FirewallInstanceResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallInstanceResource struct{}"
	}

	return strings.Join([]string{"FirewallInstanceResource", string(data)}, " ")
}
