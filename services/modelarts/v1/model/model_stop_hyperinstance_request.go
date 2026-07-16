package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopHyperinstanceRequest Request Object
type StopHyperinstanceRequest struct {

	// **参数解释**：Lite Server超节点ID。 **约束限制**：不涉及。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o StopHyperinstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopHyperinstanceRequest struct{}"
	}

	return strings.Join([]string{"StopHyperinstanceRequest", string(data)}, " ")
}
