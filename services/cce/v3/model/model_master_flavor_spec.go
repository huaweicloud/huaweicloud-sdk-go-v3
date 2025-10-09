package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MasterFlavorSpec struct {

	// **参数解释**： 控制节点规格 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 控制节点支持的可用区
	Azs *[]string `json:"azs,omitempty"`

	// **参数解释**： 控制节点所在可用区支持的故障域
	AzFaultDomains map[string][]string `json:"azFaultDomains,omitempty"`
}

func (o MasterFlavorSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterFlavorSpec struct{}"
	}

	return strings.Join([]string{"MasterFlavorSpec", string(data)}, " ")
}
