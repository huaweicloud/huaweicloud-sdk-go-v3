package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOpenSourceEnabledResponse Response Object
type ShowOpenSourceEnabledResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowOpenSourceEnabledResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// 参数解释: 查询结果。 取值范围: true：是； false：否。
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowOpenSourceEnabledResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpenSourceEnabledResponse struct{}"
	}

	return strings.Join([]string{"ShowOpenSourceEnabledResponse", string(data)}, " ")
}

type ShowOpenSourceEnabledResponseStatus struct {
	value string
}

type ShowOpenSourceEnabledResponseStatusEnum struct {
	SUCCESS ShowOpenSourceEnabledResponseStatus
	ERROR   ShowOpenSourceEnabledResponseStatus
}

func GetShowOpenSourceEnabledResponseStatusEnum() ShowOpenSourceEnabledResponseStatusEnum {
	return ShowOpenSourceEnabledResponseStatusEnum{
		SUCCESS: ShowOpenSourceEnabledResponseStatus{
			value: "success",
		},
		ERROR: ShowOpenSourceEnabledResponseStatus{
			value: "error",
		},
	}
}

func (c ShowOpenSourceEnabledResponseStatus) Value() string {
	return c.value
}

func (c ShowOpenSourceEnabledResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOpenSourceEnabledResponseStatus) UnmarshalJSON(b []byte) error {
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
