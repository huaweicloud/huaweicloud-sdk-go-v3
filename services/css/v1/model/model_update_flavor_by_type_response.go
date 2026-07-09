package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlavorByTypeResponse Response Object
type UpdateFlavorByTypeResponse struct {

	// **参数解释**： 变更订单ID，仅包周期集群返回。 **取值范围**： 不涉及
	OrderId *string `json:"orderId,omitempty"`

	// **参数解释**： 集群ID，仅包周期集群返回。 **取值范围**： 不涉及
	ClusterId *string `json:"clusterId,omitempty"`

	// **参数解释**： 变更模式，仅包周期集群返回。 **取值范围**： - 10：升配。 - 30：降配。
	ChangeMode     *int32 `json:"changeMode,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateFlavorByTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlavorByTypeResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlavorByTypeResponse", string(data)}, " ")
}
