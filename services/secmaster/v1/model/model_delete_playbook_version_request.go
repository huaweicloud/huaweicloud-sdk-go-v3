package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookVersionRequest Request Object
type DeletePlaybookVersionRequest struct {

	// **参数解释：** 内容类型 - application/json    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json  **默认取值：** application/json
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`
}

func (o DeletePlaybookVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookVersionRequest struct{}"
	}

	return strings.Join([]string{"DeletePlaybookVersionRequest", string(data)}, " ")
}
