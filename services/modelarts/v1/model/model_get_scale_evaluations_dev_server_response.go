package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetScaleEvaluationsDevServerResponse Response Object
type GetScaleEvaluationsDevServerResponse struct {

	// **参数解释**：规格容量保有情况 **约束限制**：不涉及 **取值范围**：不涉及。 **默认取值**：不涉及。
	Evaluations    *[]ServerScaleEvaluation `json:"evaluations,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o GetScaleEvaluationsDevServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetScaleEvaluationsDevServerResponse struct{}"
	}

	return strings.Join([]string{"GetScaleEvaluationsDevServerResponse", string(data)}, " ")
}
