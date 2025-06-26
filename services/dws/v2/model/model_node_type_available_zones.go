package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTypeAvailableZones **参数解释**： 规格支持的可用区及状态信息。 **取值范围**： 不涉及。
type NodeTypeAvailableZones struct {

	// **参数解释**： 可用区ID。 **取值范围**： 不涉及。
	Code string `json:"code"`

	// **参数解释**： 规格可用状态。 **取值范围**： - normal：可用 - sellout：售罄 - abandon：不可用
	Status string `json:"status"`
}

func (o NodeTypeAvailableZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypeAvailableZones struct{}"
	}

	return strings.Join([]string{"NodeTypeAvailableZones", string(data)}, " ")
}
