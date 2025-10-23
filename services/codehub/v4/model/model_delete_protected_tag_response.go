package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteProtectedTagResponse Response Object
type DeleteProtectedTagResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *DeleteProtectedTagResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o DeleteProtectedTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedTagResponse", string(data)}, " ")
}

type DeleteProtectedTagResponseStatus struct {
	value string
}

type DeleteProtectedTagResponseStatusEnum struct {
	SUCCESS DeleteProtectedTagResponseStatus
	FAIL    DeleteProtectedTagResponseStatus
}

func GetDeleteProtectedTagResponseStatusEnum() DeleteProtectedTagResponseStatusEnum {
	return DeleteProtectedTagResponseStatusEnum{
		SUCCESS: DeleteProtectedTagResponseStatus{
			value: "success",
		},
		FAIL: DeleteProtectedTagResponseStatus{
			value: "fail",
		},
	}
}

func (c DeleteProtectedTagResponseStatus) Value() string {
	return c.value
}

func (c DeleteProtectedTagResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteProtectedTagResponseStatus) UnmarshalJSON(b []byte) error {
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
