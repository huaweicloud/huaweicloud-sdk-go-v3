package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScheduledEventsRequest Request Object
type ListScheduledEventsRequest struct {

	// **参数解释**：工作空间ID，默认值为0，取值于查询workspaces列表的接口的id字段。 **约束限制**：系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长小于63个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkspaceId *string `json:"workspaceId,omitempty"`

	// **参数解释**：事件状态。 **约束限制**：不涉及。 **取值范围**：可选择值如下： - inquiring: 待授权, - scheduled: 待执行 - executing: 执行中 - completed: 执行成功 - failed: 执行失败 - canceled: 取消 **默认取值**：不涉及。
	State *[]ListScheduledEventsRequestState `json:"state,omitempty"`

	// **参数解释**：事件类型。 **约束限制**：不涉及。 **取值范围**：可选择值如下： - system-maintenance：系统维护 - localdisk-recovery：本地盘恢复 - node_reboot：节点重启 - operation-request：运维授权 - node_maintenance：超节点维护 - node_redeploy：超节点重部署 - node_localdisk_recovery 超节点本地盘恢复 **默认取值**：不涉及。
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：计划事件ID，取值查询计划事件列表接口的event_id字段。 **约束限制**：系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，长度小于63。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：节点名称，取值自节点详情的metadata.name字段。 **约束限制**：系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，小于63个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodeName *string `json:"nodeName,omitempty"`

	// **参数解释**：资源池名称, lite-cluster、standard才具有，取值自资源池详情的metadata.name字段。查询指定standard cluster和lite cluster下节点的计划事件时可传递该参数。 **约束限制**：系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，小于63个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName *string `json:"poolName,omitempty"`

	// **参数解释**：事件发布开始时间,按照时间范围过滤。 **约束限制**：按照ISO8601标准表示，并使用UTC +0时间，格式为yyyy-MM-ddTHH:mm:ssZ。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PublishStartTime *string `json:"publishStartTime,omitempty"`

	// **参数解释**：事件发布结束时间,按照时间范围过滤。 **约束限制**：按照ISO8601标准表示，并使用UTC +0时间，格式为yyyy-MM-ddTHH:mm:ssZ。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PublishEndTime *string `json:"publishEndTime,omitempty"`

	// **参数解释**：偏移量,表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：[0,1000000000]。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：[1,100]。 **默认取值**：100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduledEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledEventsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledEventsRequest", string(data)}, " ")
}

type ListScheduledEventsRequestState struct {
	value string
}

type ListScheduledEventsRequestStateEnum struct {
	INQUIRING ListScheduledEventsRequestState
	SCHEDULED ListScheduledEventsRequestState
	EXECUTING ListScheduledEventsRequestState
	COMPLETED ListScheduledEventsRequestState
	FAILED    ListScheduledEventsRequestState
	CANCELED  ListScheduledEventsRequestState
}

func GetListScheduledEventsRequestStateEnum() ListScheduledEventsRequestStateEnum {
	return ListScheduledEventsRequestStateEnum{
		INQUIRING: ListScheduledEventsRequestState{
			value: "inquiring",
		},
		SCHEDULED: ListScheduledEventsRequestState{
			value: "scheduled",
		},
		EXECUTING: ListScheduledEventsRequestState{
			value: "executing",
		},
		COMPLETED: ListScheduledEventsRequestState{
			value: "completed",
		},
		FAILED: ListScheduledEventsRequestState{
			value: "failed",
		},
		CANCELED: ListScheduledEventsRequestState{
			value: "canceled",
		},
	}
}

func (c ListScheduledEventsRequestState) Value() string {
	return c.value
}

func (c ListScheduledEventsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledEventsRequestState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
