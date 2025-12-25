package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModuleRequest Request Object
type CreateModuleRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *ModuleCreateRequestPojo `json:"body,omitempty"`
}

func (o CreateModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModuleRequest struct{}"
	}

	return strings.Join([]string{"CreateModuleRequest", string(data)}, " ")
}
