package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataobjectRelationRequest Request Object
type DeleteDataobjectRelationRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 关联主体数据对象所属数据类，小写复数，如告警为alerts，事件为incidents
	DataclassType string `json:"dataclass_type"`

	// 关联主体数据对象的id
	DataObjectId string `json:"data_object_id"`

	// 被关联的数据对象所属数据类，小写复数，如告警为alerts，事件为incidents
	RelatedDataclassType string `json:"related_dataclass_type"`

	Body *CreateDataobjectRelationRequestBody `json:"body,omitempty"`
}

func (o DeleteDataobjectRelationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataobjectRelationRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataobjectRelationRequest", string(data)}, " ")
}
