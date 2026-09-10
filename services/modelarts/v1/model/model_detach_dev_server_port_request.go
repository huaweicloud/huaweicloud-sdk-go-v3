package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachDevServerPortRequest Request Object
type DetachDevServerPortRequest struct {

	// **参数解释**：lite Server实例ID。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`

	// **参数解释**：要卸载的网卡ID。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	PortId string `json:"port_id"`
}

func (o DetachDevServerPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachDevServerPortRequest struct{}"
	}

	return strings.Join([]string{"DetachDevServerPortRequest", string(data)}, " ")
}
