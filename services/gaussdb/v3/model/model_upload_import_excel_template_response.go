package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadImportExcelTemplateResponse Response Object
type UploadImportExcelTemplateResponse struct {

	// **参数解释**：  Excel导入返回状态。  **取值范围**：  - true: 导入成功。 - false： 导入失败。
	Success *bool `json:"success,omitempty"`

	// **参数解释**：  已处理的行数。   **取值范围**：  不涉及。
	ProcessedRows *int32 `json:"processed_rows,omitempty"`

	// **参数解释**：  导入失败返回的错误列表。
	ErrorTables *[]ErrorTable `json:"error_tables,omitempty"`

	// **参数解释**：  Excel导入验证成功的列表。
	SuccessTables *[]SuccessTable `json:"success_tables,omitempty"`

	// **参数解释**：  Excel导入验证失败的行数。   **取值范围**：  不涉及。
	ErrorCount *int32 `json:"error_count,omitempty"`

	// **参数解释**：  Excel导入验证成功的行数。   **取值范围**：  不涉及。
	SuccessCount   *int32 `json:"success_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UploadImportExcelTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadImportExcelTemplateResponse struct{}"
	}

	return strings.Join([]string{"UploadImportExcelTemplateResponse", string(data)}, " ")
}
