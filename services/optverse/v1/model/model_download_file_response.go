package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadFileResponse Response Object
type DownloadFileResponse struct {

	// 参数解释： 对话ID。 约束限制： 不涉及 取值范围： 不涉及 默认取值： 不涉及
	DownloadUrl *string `json:"download_url,omitempty"`

	// **参数解释**：   文件内容。   **约束限制**：   不涉及   **取值范围**：   不涉及 **默认取值**：   不涉及
	Content        *string `json:"content,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadFileResponse", string(data)}, " ")
}
