package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFileTreeResponse Response Object
type ShowFileTreeResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowFileTreeResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *ShowFileTreeResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowFileTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowFileTreeResponse", string(data)}, " ")
}

type ShowFileTreeResponseStatus struct {
	value string
}

type ShowFileTreeResponseStatusEnum struct {
	SUCCESS ShowFileTreeResponseStatus
	ERROR   ShowFileTreeResponseStatus
}

func GetShowFileTreeResponseStatusEnum() ShowFileTreeResponseStatusEnum {
	return ShowFileTreeResponseStatusEnum{
		SUCCESS: ShowFileTreeResponseStatus{
			value: "success",
		},
		ERROR: ShowFileTreeResponseStatus{
			value: "error",
		},
	}
}

func (c ShowFileTreeResponseStatus) Value() string {
	return c.value
}

func (c ShowFileTreeResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFileTreeResponseStatus) UnmarshalJSON(b []byte) error {
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
