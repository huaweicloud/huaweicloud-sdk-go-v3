package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlinkSavepointRequest Request Object
type ListFlinkSavepointRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：Flink作业ID，用于唯一标识一个Flink作业，可通过[查询Flink作业列表](ListFlinkJobs.xml)接口获取。 **约束限制**：不涉及。 **取值范围**：采用UUIDv7格式，固定36个字符。 **默认取值**：不涉及。
	JobId string `json:"job_id"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之后的作业。unix时间戳，单位：毫秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTimeAfter *int64 `json:"create_time_after,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之前的作业。unix时间戳，单位：毫秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTimeBefore *int64 `json:"create_time_before,omitempty"`

	// **参数解释**：Savepoint状态，用于按状态筛选Savepoint列表。 **约束限制**：不涉及。 **取值范围**：   - TRIGGER_PENDING：排队等待执行。   - IN_PROGRESS：执行中。   - COMPLETED：触发成功。   - FAILED：触发失败。   - ABANDONED：放弃/失效。 **默认取值**：不涉及。
	State *ListFlinkSavepointRequestState `json:"state,omitempty"`

	// **参数解释**：Savepoint触发类型，用于按触发类型筛选Savepoint列表。 **约束限制**：不涉及。 **取值范围**：固定值USER_REQUEST，表示由用户手动生成。 **默认取值**：不涉及。
	TriggerType *ListFlinkSavepointRequestTriggerType `json:"trigger_type,omitempty"`
}

func (o ListFlinkSavepointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkSavepointRequest struct{}"
	}

	return strings.Join([]string{"ListFlinkSavepointRequest", string(data)}, " ")
}

type ListFlinkSavepointRequestState struct {
	value string
}

type ListFlinkSavepointRequestStateEnum struct {
	TRIGGER_PENDING ListFlinkSavepointRequestState
	IN_PROGRESS     ListFlinkSavepointRequestState
	COMPLETED       ListFlinkSavepointRequestState
	FAILED          ListFlinkSavepointRequestState
	ABANDONED       ListFlinkSavepointRequestState
}

func GetListFlinkSavepointRequestStateEnum() ListFlinkSavepointRequestStateEnum {
	return ListFlinkSavepointRequestStateEnum{
		TRIGGER_PENDING: ListFlinkSavepointRequestState{
			value: "TRIGGER_PENDING",
		},
		IN_PROGRESS: ListFlinkSavepointRequestState{
			value: "IN_PROGRESS",
		},
		COMPLETED: ListFlinkSavepointRequestState{
			value: "COMPLETED",
		},
		FAILED: ListFlinkSavepointRequestState{
			value: "FAILED",
		},
		ABANDONED: ListFlinkSavepointRequestState{
			value: "ABANDONED",
		},
	}
}

func (c ListFlinkSavepointRequestState) Value() string {
	return c.value
}

func (c ListFlinkSavepointRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlinkSavepointRequestState) UnmarshalJSON(b []byte) error {
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

type ListFlinkSavepointRequestTriggerType struct {
	value string
}

type ListFlinkSavepointRequestTriggerTypeEnum struct {
	USER_REQUEST ListFlinkSavepointRequestTriggerType
}

func GetListFlinkSavepointRequestTriggerTypeEnum() ListFlinkSavepointRequestTriggerTypeEnum {
	return ListFlinkSavepointRequestTriggerTypeEnum{
		USER_REQUEST: ListFlinkSavepointRequestTriggerType{
			value: "USER_REQUEST",
		},
	}
}

func (c ListFlinkSavepointRequestTriggerType) Value() string {
	return c.value
}

func (c ListFlinkSavepointRequestTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlinkSavepointRequestTriggerType) UnmarshalJSON(b []byte) error {
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
