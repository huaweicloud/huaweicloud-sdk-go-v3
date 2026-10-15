package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSparkOperatorStateResponse Response Object
type ShowSparkOperatorStateResponse struct {

	// **参数解释**：操作状态，用于标识异步操作的当前执行状态。 **取值范围**：   - RUNNING：运行中。   - FAILED：失败。   - SUCCESS：成功。
	Status *ShowSparkOperatorStateResponseStatus `json:"status,omitempty"`

	OperatorInfo *ShowSparkOperatorStateOperationInfo `json:"operator_info,omitempty"`

	// **参数解释**：操作描述信息，包含操作状态描述或失败时的错误信息。 **取值范围**：长度为0~1024个字符。
	Message *string `json:"message,omitempty"`

	// **参数解释**：操作创建时间，用于记录操作提交的时间。 **取值范围**：Unix时间戳，单位为毫秒，取值范围为0~9223372036854775807。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**：操作结束时间，用于记录操作完成的时间。 **取值范围**：Unix时间戳，单位为毫秒，取值范围为0~9223372036854775807。
	EndTime        *int64 `json:"end_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSparkOperatorStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkOperatorStateResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkOperatorStateResponse", string(data)}, " ")
}

type ShowSparkOperatorStateResponseStatus struct {
	value string
}

type ShowSparkOperatorStateResponseStatusEnum struct {
	RUNNING ShowSparkOperatorStateResponseStatus
	FAILED  ShowSparkOperatorStateResponseStatus
	SUCCESS ShowSparkOperatorStateResponseStatus
}

func GetShowSparkOperatorStateResponseStatusEnum() ShowSparkOperatorStateResponseStatusEnum {
	return ShowSparkOperatorStateResponseStatusEnum{
		RUNNING: ShowSparkOperatorStateResponseStatus{
			value: "RUNNING",
		},
		FAILED: ShowSparkOperatorStateResponseStatus{
			value: "FAILED",
		},
		SUCCESS: ShowSparkOperatorStateResponseStatus{
			value: "SUCCESS",
		},
	}
}

func (c ShowSparkOperatorStateResponseStatus) Value() string {
	return c.value
}

func (c ShowSparkOperatorStateResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSparkOperatorStateResponseStatus) UnmarshalJSON(b []byte) error {
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
