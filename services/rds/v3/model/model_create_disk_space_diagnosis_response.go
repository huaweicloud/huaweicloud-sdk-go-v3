package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiskSpaceDiagnosisResponse Response Object
type CreateDiskSpaceDiagnosisResponse struct {

	// **参数解释**：   下发结果。  **约束限制**：   不涉及。  **取值范围**：  - success 代表后台开始诊断。  **默认取值**：   不涉及。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDiskSpaceDiagnosisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiskSpaceDiagnosisResponse struct{}"
	}

	return strings.Join([]string{"CreateDiskSpaceDiagnosisResponse", string(data)}, " ")
}
