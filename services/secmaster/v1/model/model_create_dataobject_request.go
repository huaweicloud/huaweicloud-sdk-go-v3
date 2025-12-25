package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataobjectRequest Request Object
type CreateDataobjectRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 关联主体数据对象所属数据类名称，小写复数，如告警为alerts，事件为incidents
	DataclassName string `json:"dataclass_name"`

	Body *DataObjectCreateForm `json:"body,omitempty"`
}

func (o CreateDataobjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataobjectRequest struct{}"
	}

	return strings.Join([]string{"CreateDataobjectRequest", string(data)}, " ")
}
