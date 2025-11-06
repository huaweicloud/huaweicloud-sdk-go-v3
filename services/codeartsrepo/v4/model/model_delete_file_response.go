package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteFileResponse Response Object
type DeleteFileResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *DeleteFileResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o DeleteFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteFileResponse", string(data)}, " ")
}

type DeleteFileResponseStatus struct {
	value string
}

type DeleteFileResponseStatusEnum struct {
	SUCCESS DeleteFileResponseStatus
	FAIL    DeleteFileResponseStatus
}

func GetDeleteFileResponseStatusEnum() DeleteFileResponseStatusEnum {
	return DeleteFileResponseStatusEnum{
		SUCCESS: DeleteFileResponseStatus{
			value: "success",
		},
		FAIL: DeleteFileResponseStatus{
			value: "fail",
		},
	}
}

func (c DeleteFileResponseStatus) Value() string {
	return c.value
}

func (c DeleteFileResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteFileResponseStatus) UnmarshalJSON(b []byte) error {
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
