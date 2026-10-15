package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAuraSessionStatementRecordsRequest Request Object
type ListAuraSessionStatementRecordsRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：会话ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	SessionId string `json:"session_id"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：上一页中最后一条记录id，查询第一页时传空值。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：statement id。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：状态过滤，支持一种状态查询，默认查询所有。 **约束限制**：不涉及。 **取值范围**：   - CANCELED：已取消。   - FAILED：失败。   - SUCCESSFUL：成功。   - RUNNING：运行中。   - SUBMITTED：已提交。   - ERROR：错误。 **默认取值**：不涉及。
	Status *ListAuraSessionStatementRecordsRequestStatus `json:"status,omitempty"`
}

func (o ListAuraSessionStatementRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraSessionStatementRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListAuraSessionStatementRecordsRequest", string(data)}, " ")
}

type ListAuraSessionStatementRecordsRequestStatus struct {
	value string
}

type ListAuraSessionStatementRecordsRequestStatusEnum struct {
	CANCELED   ListAuraSessionStatementRecordsRequestStatus
	FAILED     ListAuraSessionStatementRecordsRequestStatus
	SUCCESSFUL ListAuraSessionStatementRecordsRequestStatus
	RUNNING    ListAuraSessionStatementRecordsRequestStatus
	SUBMITTED  ListAuraSessionStatementRecordsRequestStatus
	ERROR      ListAuraSessionStatementRecordsRequestStatus
}

func GetListAuraSessionStatementRecordsRequestStatusEnum() ListAuraSessionStatementRecordsRequestStatusEnum {
	return ListAuraSessionStatementRecordsRequestStatusEnum{
		CANCELED: ListAuraSessionStatementRecordsRequestStatus{
			value: "CANCELED",
		},
		FAILED: ListAuraSessionStatementRecordsRequestStatus{
			value: "FAILED",
		},
		SUCCESSFUL: ListAuraSessionStatementRecordsRequestStatus{
			value: "SUCCESSFUL",
		},
		RUNNING: ListAuraSessionStatementRecordsRequestStatus{
			value: "RUNNING",
		},
		SUBMITTED: ListAuraSessionStatementRecordsRequestStatus{
			value: "SUBMITTED",
		},
		ERROR: ListAuraSessionStatementRecordsRequestStatus{
			value: "ERROR",
		},
	}
}

func (c ListAuraSessionStatementRecordsRequestStatus) Value() string {
	return c.value
}

func (c ListAuraSessionStatementRecordsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAuraSessionStatementRecordsRequestStatus) UnmarshalJSON(b []byte) error {
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
