package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisStackRequest Request Object
type ShowDiagnosisStackRequest struct {

	// **参数解释**： 引擎。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 堆ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StackId string `json:"stack_id"`
}

func (o ShowDiagnosisStackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisStackRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisStackRequest", string(data)}, " ")
}
