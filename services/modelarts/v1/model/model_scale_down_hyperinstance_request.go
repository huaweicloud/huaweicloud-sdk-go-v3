package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleDownHyperinstanceRequest Request Object
type ScaleDownHyperinstanceRequest struct {

	// **参数解释**：Lite Server超节点ID。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`

	Body *ServerScaleDownRequest `json:"body,omitempty"`
}

func (o ScaleDownHyperinstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleDownHyperinstanceRequest struct{}"
	}

	return strings.Join([]string{"ScaleDownHyperinstanceRequest", string(data)}, " ")
}
