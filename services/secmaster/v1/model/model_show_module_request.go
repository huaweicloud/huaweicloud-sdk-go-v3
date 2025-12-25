package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModuleRequest Request Object
type ShowModuleRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 模板ID
	ModuleId string `json:"module_id"`
}

func (o ShowModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModuleRequest struct{}"
	}

	return strings.Join([]string{"ShowModuleRequest", string(data)}, " ")
}
