package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteLayoutRequest Request Object
type ExecuteLayoutRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *ExecuteLayoutRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ExecuteLayoutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteLayoutRequest struct{}"
	}

	return strings.Join([]string{"ExecuteLayoutRequest", string(data)}, " ")
}
