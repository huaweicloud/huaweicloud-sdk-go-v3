package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataclassTypeRequest Request Object
type DeleteDataclassTypeRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 数据类的唯一ID
	DataclassId string `json:"dataclass_id"`

	Body *DeleteDataclassTypeRequestBody `json:"body,omitempty"`
}

func (o DeleteDataclassTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataclassTypeRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataclassTypeRequest", string(data)}, " ")
}
