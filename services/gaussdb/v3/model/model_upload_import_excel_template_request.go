package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadImportExcelTemplateRequest Request Object
type UploadImportExcelTemplateRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  HTAP标准版实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *UploadImportExcelTemplateRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadImportExcelTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadImportExcelTemplateRequest struct{}"
	}

	return strings.Join([]string{"UploadImportExcelTemplateRequest", string(data)}, " ")
}
