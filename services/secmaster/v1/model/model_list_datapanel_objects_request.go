package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatapanelObjectsRequest Request Object
type ListDatapanelObjectsRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 数据类ID, 您可通过调用[查询数据类列表](ListDataclass.xml)接口获取数据类ID
	Dataclass string `json:"dataclass"`

	Body *DataobjectSearch `json:"body,omitempty"`
}

func (o ListDatapanelObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatapanelObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListDatapanelObjectsRequest", string(data)}, " ")
}
