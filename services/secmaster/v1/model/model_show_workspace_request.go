package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceRequest Request Object
type ShowWorkspaceRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ShowWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceRequest", string(data)}, " ")
}
