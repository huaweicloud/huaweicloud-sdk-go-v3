package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMavenUserResponse Response Object
type ListMavenUserResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ListMavenUserResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *DevcloudUserDov5Page `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListMavenUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMavenUserResponse struct{}"
	}

	return strings.Join([]string{"ListMavenUserResponse", string(data)}, " ")
}

type ListMavenUserResponseStatus struct {
	value string
}

type ListMavenUserResponseStatusEnum struct {
	SUCCESS ListMavenUserResponseStatus
	ERROR   ListMavenUserResponseStatus
}

func GetListMavenUserResponseStatusEnum() ListMavenUserResponseStatusEnum {
	return ListMavenUserResponseStatusEnum{
		SUCCESS: ListMavenUserResponseStatus{
			value: "success",
		},
		ERROR: ListMavenUserResponseStatus{
			value: "error",
		},
	}
}

func (c ListMavenUserResponseStatus) Value() string {
	return c.value
}

func (c ListMavenUserResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMavenUserResponseStatus) UnmarshalJSON(b []byte) error {
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
