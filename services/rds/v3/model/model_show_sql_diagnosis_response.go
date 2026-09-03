package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlDiagnosisResponse Response Object
type ShowSqlDiagnosisResponse struct {

	// **参数解释**：  sql信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Results *[]SqlDiagnosisResult `json:"results,omitempty"`

	// **参数解释**：  总数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSqlDiagnosisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlDiagnosisResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlDiagnosisResponse", string(data)}, " ")
}
