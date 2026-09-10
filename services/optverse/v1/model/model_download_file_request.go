package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadFileRequest Request Object
type DownloadFileRequest struct {

	// **参数解释**：   返回文件内容。   **约束限制**：   不涉及 **取值范围**：   * true：返回文件内容 * false：不返回文件内容 **默认取值**：   false
	XNeedContent *bool `json:"X-Need-Content,omitempty"`

	// **参数解释**： 对话ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	ChatId string `json:"chat_id"`

	// **参数解释**： 问答ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	Filename string `json:"filename"`
}

func (o DownloadFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFileRequest struct{}"
	}

	return strings.Join([]string{"DownloadFileRequest", string(data)}, " ")
}
