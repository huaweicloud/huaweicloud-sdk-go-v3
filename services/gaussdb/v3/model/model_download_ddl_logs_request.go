package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDdlLogsRequest Request Object
type DownloadDdlLogsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  租户在某一project下的实例id  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字、下划线组成，且长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *DownloadDdlLogsRequestBody `json:"body,omitempty"`
}

func (o DownloadDdlLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDdlLogsRequest struct{}"
	}

	return strings.Join([]string{"DownloadDdlLogsRequest", string(data)}, " ")
}
