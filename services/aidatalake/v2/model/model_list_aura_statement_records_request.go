package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAuraStatementRecordsRequest Request Object
type ListAuraStatementRecordsRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：通过ID检索Endpoint的参数。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：通过name检索Endpoint的参数。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	// **参数解释**：Session的ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：statement id。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：执行sql来源。 **约束限制**：不涉及。 **取值范围**：   - 0：API。   - 1：Console。 **默认取值**：不涉及。
	Source *int32 `json:"source,omitempty"`

	// **参数解释**：状态过滤，支持一种状态查询，默认查询所有。 **约束限制**：不涉及。 **取值范围**：   - CANCELED：已取消。   - FAILED：失败。   - SUCCESSFUL：成功。   - RUNNING：运行中。   - SUBMITTED：已提交。   - ERROR：错误。 **默认取值**：不涉及。
	Status *ListAuraStatementRecordsRequestStatus `json:"status,omitempty"`

	// **参数解释**： statement记录列表查询起始时间，该时间为statement创建时间，时间戳，单位：秒。 **约束限制**：不涉及。 **取值范围**：1~2147483647。 **默认取值**：不涉及。
	StartTime *int32 `json:"start_time,omitempty"`

	// **参数解释**：statement记录列表查询结束时间，该时间为statement创建时间，时间戳，单位：秒。 **约束限制**：不涉及。 **取值范围**：1~2147483647。 **默认取值**：不涉及。
	EndTime *int32 `json:"end_time,omitempty"`

	// **参数解释**：上一页中最后一条记录id，查询第一页时传空值。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAuraStatementRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraStatementRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListAuraStatementRecordsRequest", string(data)}, " ")
}

type ListAuraStatementRecordsRequestStatus struct {
	value string
}

type ListAuraStatementRecordsRequestStatusEnum struct {
	CANCELED   ListAuraStatementRecordsRequestStatus
	FAILED     ListAuraStatementRecordsRequestStatus
	SUCCESSFUL ListAuraStatementRecordsRequestStatus
	RUNNING    ListAuraStatementRecordsRequestStatus
	SUBMITTED  ListAuraStatementRecordsRequestStatus
	ERROR      ListAuraStatementRecordsRequestStatus
}

func GetListAuraStatementRecordsRequestStatusEnum() ListAuraStatementRecordsRequestStatusEnum {
	return ListAuraStatementRecordsRequestStatusEnum{
		CANCELED: ListAuraStatementRecordsRequestStatus{
			value: "CANCELED",
		},
		FAILED: ListAuraStatementRecordsRequestStatus{
			value: "FAILED",
		},
		SUCCESSFUL: ListAuraStatementRecordsRequestStatus{
			value: "SUCCESSFUL",
		},
		RUNNING: ListAuraStatementRecordsRequestStatus{
			value: "RUNNING",
		},
		SUBMITTED: ListAuraStatementRecordsRequestStatus{
			value: "SUBMITTED",
		},
		ERROR: ListAuraStatementRecordsRequestStatus{
			value: "ERROR",
		},
	}
}

func (c ListAuraStatementRecordsRequestStatus) Value() string {
	return c.value
}

func (c ListAuraStatementRecordsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAuraStatementRecordsRequestStatus) UnmarshalJSON(b []byte) error {
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
