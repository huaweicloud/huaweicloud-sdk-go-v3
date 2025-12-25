package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFileDetailByFullNameResponse Response Object
type ShowFileDetailByFullNameResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowFileDetailByFullNameResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *RepoFileDov5 `json:"result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowFileDetailByFullNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileDetailByFullNameResponse struct{}"
	}

	return strings.Join([]string{"ShowFileDetailByFullNameResponse", string(data)}, " ")
}

type ShowFileDetailByFullNameResponseStatus struct {
	value string
}

type ShowFileDetailByFullNameResponseStatusEnum struct {
	SUCCESS ShowFileDetailByFullNameResponseStatus
	ERROR   ShowFileDetailByFullNameResponseStatus
}

func GetShowFileDetailByFullNameResponseStatusEnum() ShowFileDetailByFullNameResponseStatusEnum {
	return ShowFileDetailByFullNameResponseStatusEnum{
		SUCCESS: ShowFileDetailByFullNameResponseStatus{
			value: "success",
		},
		ERROR: ShowFileDetailByFullNameResponseStatus{
			value: "error",
		},
	}
}

func (c ShowFileDetailByFullNameResponseStatus) Value() string {
	return c.value
}

func (c ShowFileDetailByFullNameResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFileDetailByFullNameResponseStatus) UnmarshalJSON(b []byte) error {
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
