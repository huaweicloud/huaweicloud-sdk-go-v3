package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatementExecuteRequestBody **参数解释**：执行查询请求体。 **约束限制**：不涉及。
type StatementExecuteRequestBody struct {

	// **参数解释**：Session的ID。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
	SessionId string `json:"session_id"`

	// **参数解释**：SQL statement信息。 **约束限制**：不涉及。 **取值范围**：长度为1~10000000个字符。 **默认取值**：不涉及。
	Statement string `json:"statement"`

	// **参数解释**：返回结果集的最大行数限制。 **约束限制**：不涉及。 **取值范围**：1~2147483647，单位：行。 **默认取值**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：SQL语句的参数绑定列表，用于绑定变量传参。外层数组为绑定集合（实际使用首个元素 bindings[0]），内层数组为单条SQL的参数值列表，每个字符串元素为单个参数值。 **约束限制**：   - 外层数组元素个数：0~1。   - 内层数组元素个数：0~200。   - 每个参数值字符串长度：0~500个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Bindings *[][]string `json:"bindings,omitempty"`

	// **参数解释**：同步执行或异步执行。 **约束限制**：不涉及。 **取值范围**：   - true：同步。   - false：异步。 **默认取值**：true（未传时默认同步执行）。
	IsSync *bool `json:"is_sync,omitempty"`

	// **参数解释**：是否来自运维通道。 **约束限制**：不涉及。 **取值范围**：  - true：运维通道语句。  - false：不是运维通道语句。 **默认取值**：false
	IsMaintainStatement *bool `json:"is_maintain_statement,omitempty"`

	// **参数解释**：获取结果集的方式。 **约束限制**：不涉及。 **取值范围**：   - DEFAULT：直接返回结果集。   - READ_OBS：返回结果集在桶上的路径，而不是直接返回结果集。 **默认取值**：DEFAULT。
	ResultsetFetchMode *StatementExecuteRequestBodyResultsetFetchMode `json:"resultset_fetch_mode,omitempty"`

	// **参数解释**：SQL排队等待时间，单位：秒。 **约束限制**：不涉及。 **取值范围**：300~86400。 **默认取值**：不涉及。
	WaitTimeout *int32 `json:"wait_timeout,omitempty"`

	// **参数解释**：资源不足时重试次数。 **约束限制**：不涉及。 **取值范围**：1~6000。 **默认取值**：不涉及。
	MaxRetryTimes *int32 `json:"max_retry_times,omitempty"`

	// **参数解释**：作业类型。 **约束限制**：不涉及。 **取值范围**：   - SQL：SQL查询。   - JOB：作业。 **默认取值**：SQL。
	StatementType *string `json:"statement_type,omitempty"`

	// **参数解释**：作业名称。 **约束限制**：长度不超过64个字符，仅允许字母、数字、中文、下划线和连字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	StatementName *string `json:"statement_name,omitempty"`
}

func (o StatementExecuteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatementExecuteRequestBody struct{}"
	}

	return strings.Join([]string{"StatementExecuteRequestBody", string(data)}, " ")
}

type StatementExecuteRequestBodyResultsetFetchMode struct {
	value string
}

type StatementExecuteRequestBodyResultsetFetchModeEnum struct {
	DEFAULT  StatementExecuteRequestBodyResultsetFetchMode
	READ_OBS StatementExecuteRequestBodyResultsetFetchMode
}

func GetStatementExecuteRequestBodyResultsetFetchModeEnum() StatementExecuteRequestBodyResultsetFetchModeEnum {
	return StatementExecuteRequestBodyResultsetFetchModeEnum{
		DEFAULT: StatementExecuteRequestBodyResultsetFetchMode{
			value: "DEFAULT",
		},
		READ_OBS: StatementExecuteRequestBodyResultsetFetchMode{
			value: "READ_OBS",
		},
	}
}

func (c StatementExecuteRequestBodyResultsetFetchMode) Value() string {
	return c.value
}

func (c StatementExecuteRequestBodyResultsetFetchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatementExecuteRequestBodyResultsetFetchMode) UnmarshalJSON(b []byte) error {
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
