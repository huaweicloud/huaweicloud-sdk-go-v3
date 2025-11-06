package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeFlavor struct {

	// **参数解释**： 规格id。 **取值范围**： 不涉及
	StrId *string `json:"str_id,omitempty"`

	// **参数解释**： 实例的CPU核数。 **取值范围**： 不涉及
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**： 实例的内存大小。单位GB。 **取值范围**： 不涉及
	Ram *int32 `json:"ram,omitempty"`

	// **参数解释**： 规格名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 节点规格支持的Region。 **取值范围**： 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**： 实例的硬盘容量范围，单位GB。 **取值范围**： 不涉及
	Diskrange *string `json:"diskrange,omitempty"`

	// **参数解释**： 节点类型。 **取值范围**： - ess：数据节点。 - ess-cold：冷数据节点。 - ess-client：Client节点。 - ess-master：Master节点。 - lgs：LogStash节点。
	Typename *string `json:"typename,omitempty"`

	// **参数解释**： 规格售卖状态。 **取值范围**： - normal：正常商用。 - sellout：售罄。
	CondOperationStatus *string `json:"condOperationStatus,omitempty"`

	// **参数解释**： 是否本地盘。 **取值范围**： - false：非本地盘规格。 - true：本地盘规格。
	Localdisk *string `json:"localdisk,omitempty"`

	// **参数解释**： 是否边缘区规格。 **取值范围**： - false：中心可用区规格。 - true：边缘可用区规格。
	Edge *bool `json:"edge,omitempty"`
}

func (o ResizeFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeFlavor struct{}"
	}

	return strings.Join([]string{"ResizeFlavor", string(data)}, " ")
}
