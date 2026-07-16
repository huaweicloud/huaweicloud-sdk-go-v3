package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptScheduledEventRequest Request Object
type AcceptScheduledEventRequest struct {

	// **参数解释**：计划事件ID，取值查询计划事件列表接口的event_id字段。 **约束限制**：不涉及。 **取值范围**：系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，长度小于63 **默认取值**：不涉及。
	EventId string `json:"event_id"`

	// **参数解释**：工作空间ID，默认值为0，取值于查询workspaces列表的接口的id字段。 **约束限制**：系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度小于63个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkspaceId *string `json:"workspaceId,omitempty"`

	Body *EventUpdate `json:"body,omitempty"`
}

func (o AcceptScheduledEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptScheduledEventRequest struct{}"
	}

	return strings.Join([]string{"AcceptScheduledEventRequest", string(data)}, " ")
}
