package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiskSpaceDiagnosisResponse Response Object
type ShowDiskSpaceDiagnosisResponse struct {

	// **参数解释**：  诊断结果执行状态。  **约束限制**：  不涉及。  **取值范围**：  -FINISHED (已完成) -RUNNING (诊断中) -UNEXECUTED (未执行诊断)  **默认取值**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  各维度诊断信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Results        *[]DiskSpaceDiagnosisResult `json:"results,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowDiskSpaceDiagnosisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiskSpaceDiagnosisResponse struct{}"
	}

	return strings.Join([]string{"ShowDiskSpaceDiagnosisResponse", string(data)}, " ")
}
