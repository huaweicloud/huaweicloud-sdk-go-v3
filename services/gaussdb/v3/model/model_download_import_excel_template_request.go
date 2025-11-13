package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadImportExcelTemplateRequest Request Object
type DownloadImportExcelTemplateRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  HTAP标准版实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  具体选择哪一种模板进行下载。  **约束限制**：  不涉及。  **取值范围**：  import_async: Excel导入文件类型。  **默认取值**：  不涉及。
	TemplateType string `json:"template_type"`
}

func (o DownloadImportExcelTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadImportExcelTemplateRequest struct{}"
	}

	return strings.Join([]string{"DownloadImportExcelTemplateRequest", string(data)}, " ")
}
