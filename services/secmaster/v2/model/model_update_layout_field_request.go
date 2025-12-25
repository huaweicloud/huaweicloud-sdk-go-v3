package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLayoutFieldRequest Request Object
type UpdateLayoutFieldRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 字段ID
	FieldId string `json:"field_id"`

	Body *UpdateLayoutFieldRequestBody `json:"body,omitempty"`
}

func (o UpdateLayoutFieldRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLayoutFieldRequest struct{}"
	}

	return strings.Join([]string{"UpdateLayoutFieldRequest", string(data)}, " ")
}
