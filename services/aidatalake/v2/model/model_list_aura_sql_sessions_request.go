package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAuraSqlSessionsRequest Request Object
type ListAuraSqlSessionsRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：通过ID检索Endpoint的参数。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：通过name检索Endpoint的参数。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	// **参数解释**：会话ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：状态过滤，支持一种状态查询，默认查询所有。 **约束限制**：不涉及。 **取值范围**：   - RUNNING：运行中。   - CLOSED：已关闭。   - WAITING：等待中。   - CREATING：创建中。   - FAIL：失败。 **默认取值**：不涉及。
	Status *ListAuraSqlSessionsRequestStatus `json:"status,omitempty"`

	// **参数解释**： statement记录列表查询起始时间，该时间为statement创建时间，时间戳，单位：秒。 **约束限制**：不涉及。 **取值范围**：1~2147483647。 **默认取值**：不涉及。
	StartTime *int32 `json:"start_time,omitempty"`

	// **参数解释**：statement记录列表查询结束时间，该时间为statement创建时间，时间戳，单位：秒。 **约束限制**：不涉及。 **取值范围**：1~2147483647。 **默认取值**：不涉及。
	EndTime *int32 `json:"end_time,omitempty"`

	// **参数解释**：上一页中最后一条记录id，查询第一页时传空值。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAuraSqlSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraSqlSessionsRequest struct{}"
	}

	return strings.Join([]string{"ListAuraSqlSessionsRequest", string(data)}, " ")
}

type ListAuraSqlSessionsRequestStatus struct {
	value string
}

type ListAuraSqlSessionsRequestStatusEnum struct {
	RUNNING  ListAuraSqlSessionsRequestStatus
	CLOSED   ListAuraSqlSessionsRequestStatus
	WAITING  ListAuraSqlSessionsRequestStatus
	CREATING ListAuraSqlSessionsRequestStatus
	FAIL     ListAuraSqlSessionsRequestStatus
}

func GetListAuraSqlSessionsRequestStatusEnum() ListAuraSqlSessionsRequestStatusEnum {
	return ListAuraSqlSessionsRequestStatusEnum{
		RUNNING: ListAuraSqlSessionsRequestStatus{
			value: "RUNNING",
		},
		CLOSED: ListAuraSqlSessionsRequestStatus{
			value: "CLOSED",
		},
		WAITING: ListAuraSqlSessionsRequestStatus{
			value: "WAITING",
		},
		CREATING: ListAuraSqlSessionsRequestStatus{
			value: "CREATING",
		},
		FAIL: ListAuraSqlSessionsRequestStatus{
			value: "FAIL",
		},
	}
}

func (c ListAuraSqlSessionsRequestStatus) Value() string {
	return c.value
}

func (c ListAuraSqlSessionsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAuraSqlSessionsRequestStatus) UnmarshalJSON(b []byte) error {
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
