package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type SessionStatementRecord struct {

	// **参数解释**：statement id。 **取值范围**：长度为1~36的英文字母、数字、连字符和下划线的组合。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：查询语句。 **取值范围**：不涉及。
	StatementContent *string `json:"statement_content,omitempty"`

	// **参数解释**：创建时间，时间戳，单位：毫秒。 **取值范围**：不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：状态。 **取值范围**：    - CANCELED：取消。   - FAILED：失败。   - SUCCESSFUL：成功。   - RUNNING：运行中。   - SUBMITTED：提交。   - ERROR：错误。
	Status *SessionStatementRecordStatus `json:"status,omitempty"`

	// **参数解释**：用户名称。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、下划线。
	UserName *string `json:"user_name,omitempty"`
}

func (o SessionStatementRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionStatementRecord struct{}"
	}

	return strings.Join([]string{"SessionStatementRecord", string(data)}, " ")
}

type SessionStatementRecordStatus struct {
	value string
}

type SessionStatementRecordStatusEnum struct {
	CANCELED   SessionStatementRecordStatus
	FAILED     SessionStatementRecordStatus
	SUCCESSFUL SessionStatementRecordStatus
	RUNNING    SessionStatementRecordStatus
	SUBMITTED  SessionStatementRecordStatus
	ERROR      SessionStatementRecordStatus
}

func GetSessionStatementRecordStatusEnum() SessionStatementRecordStatusEnum {
	return SessionStatementRecordStatusEnum{
		CANCELED: SessionStatementRecordStatus{
			value: "CANCELED",
		},
		FAILED: SessionStatementRecordStatus{
			value: "FAILED",
		},
		SUCCESSFUL: SessionStatementRecordStatus{
			value: "SUCCESSFUL",
		},
		RUNNING: SessionStatementRecordStatus{
			value: "RUNNING",
		},
		SUBMITTED: SessionStatementRecordStatus{
			value: "SUBMITTED",
		},
		ERROR: SessionStatementRecordStatus{
			value: "ERROR",
		},
	}
}

func (c SessionStatementRecordStatus) Value() string {
	return c.value
}

func (c SessionStatementRecordStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SessionStatementRecordStatus) UnmarshalJSON(b []byte) error {
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
