package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataobjectRequest Request Object
type UpdateDataobjectRequest struct {

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 关联主体数据对象所属数据类名称，小写复数，如告警为alerts，事件为incidents
	DataclassName string `json:"dataclass_name"`

	// 关联主体数据对象的id
	DataObjectId string `json:"data_object_id"`

	Body *DataObjectBatchUpdateForm `json:"body,omitempty"`
}

func (o UpdateDataobjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataobjectRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataobjectRequest", string(data)}, " ")
}
