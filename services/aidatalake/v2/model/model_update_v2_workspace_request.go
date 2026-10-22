package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2WorkspaceRequest Request Object
type UpdateV2WorkspaceRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateV2WorkspaceRequestInput `json:"body,omitempty"`
}

func (o UpdateV2WorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2WorkspaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateV2WorkspaceRequest", string(data)}, " ")
}
