package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlDiagnosisResult 诊断结果
type SqlDiagnosisResult struct {

	// **参数解释**：  线程id。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id *int64 `json:"id,omitempty"`

	// **参数解释**：  用户名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：  用户host信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**：  数据库名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Db *string `json:"db,omitempty"`

	// **参数解释**：  执行开始时间  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：  sql语句。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Sql *string `json:"sql,omitempty"`
}

func (o SqlDiagnosisResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlDiagnosisResult struct{}"
	}

	return strings.Join([]string{"SqlDiagnosisResult", string(data)}, " ")
}
