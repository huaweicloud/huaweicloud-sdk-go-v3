package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// StatementRecord **参数解释**：Sql执行历史记录。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type StatementRecord struct {

	// **参数解释**：Sql执行StatementID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：工作空间ID。 **取值范围**：长度为1~36的英文字母、数字和中划线的组合。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：端点ID。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：Session的ID。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：数据库名称。 **取值范围**：长度为1~256个字符，支持大小写英文字母、数字、下划线。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**：sql脚本内容。 **取值范围**：长度为1~256个字符，支持大小写英文字母、数字、下划线。
	StatementContent *string `json:"statement_content,omitempty"`

	// **参数解释**：dpu时。 **取值范围**：不涉及。
	DpuDuration *string `json:"dpu_duration,omitempty"`

	// **参数解释**：执行时长。 **取值范围**：不涉及。
	Duration *string `json:"duration,omitempty"`

	// **参数解释**：费用。 **取值范围**：不涉及。
	DpuCost *string `json:"dpu_cost,omitempty"`

	// **参数解释**：Sql执行状态。 **取值范围**：   - CANCELED：取消。   - FAILED：失败。   - SUCCESSFUL：成功。   - RUNNING：运行中。   - SUBMITTED：提交。   - ERROR：错误。
	Status *StatementRecordStatus `json:"status,omitempty"`

	// **参数解释**：用户名称。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、下划线。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：sql创建时间，用于记录sql语句创建时间。 **取值范围**：unix时间戳，单位为毫秒。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：sql执行开始时间，时间戳，单位：毫秒。 **取值范围**：不涉及。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// **参数解释**：sql执行结束时间，时间戳，单位：毫秒。 **取值范围**：不涉及。
	FinishTime *sdktime.SdkTime `json:"finish_time,omitempty"`
}

func (o StatementRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatementRecord struct{}"
	}

	return strings.Join([]string{"StatementRecord", string(data)}, " ")
}

type StatementRecordStatus struct {
	value string
}

type StatementRecordStatusEnum struct {
	CANCELED   StatementRecordStatus
	FAILED     StatementRecordStatus
	SUCCESSFUL StatementRecordStatus
	RUNNING    StatementRecordStatus
	SUBMITTED  StatementRecordStatus
	ERROR      StatementRecordStatus
}

func GetStatementRecordStatusEnum() StatementRecordStatusEnum {
	return StatementRecordStatusEnum{
		CANCELED: StatementRecordStatus{
			value: "CANCELED",
		},
		FAILED: StatementRecordStatus{
			value: "FAILED",
		},
		SUCCESSFUL: StatementRecordStatus{
			value: "SUCCESSFUL",
		},
		RUNNING: StatementRecordStatus{
			value: "RUNNING",
		},
		SUBMITTED: StatementRecordStatus{
			value: "SUBMITTED",
		},
		ERROR: StatementRecordStatus{
			value: "ERROR",
		},
	}
}

func (c StatementRecordStatus) Value() string {
	return c.value
}

func (c StatementRecordStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatementRecordStatus) UnmarshalJSON(b []byte) error {
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
