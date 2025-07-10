package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PanguDrugModelResourceRsp struct {

	// **参数解释**： 实例ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id string `json:"id"`

	// **参数解释**： 资源ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ResourceId string `json:"resource_id"`

	// **参数解释**： 规格信息。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SpecCode string `json:"spec_code"`

	// **参数解释**： 计费类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ChargeMode string `json:"charge_mode"`

	// **参数解释**： 项目ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ProjectId string `json:"project_id"`

	Status *DrugModelResourceStatusEnum `json:"status"`

	// **参数解释**： 购买时间，UTC时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime string `json:"create_time"`
}

func (o PanguDrugModelResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PanguDrugModelResourceRsp struct{}"
	}

	return strings.Join([]string{"PanguDrugModelResourceRsp", string(data)}, " ")
}
