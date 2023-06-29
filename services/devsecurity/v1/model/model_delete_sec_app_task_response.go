package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSecAppTaskResponse Response Object
type DeleteSecAppTaskResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *DeleteSecAppTaskResponseInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o DeleteSecAppTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecAppTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecAppTaskResponse", string(data)}, " ")
}

type DeleteSecAppTaskResponseInfoCode struct {
	value string
}

type DeleteSecAppTaskResponseInfoCodeEnum struct {
	SUCCESS DeleteSecAppTaskResponseInfoCode
	FAILURE DeleteSecAppTaskResponseInfoCode
}

func GetDeleteSecAppTaskResponseInfoCodeEnum() DeleteSecAppTaskResponseInfoCodeEnum {
	return DeleteSecAppTaskResponseInfoCodeEnum{
		SUCCESS: DeleteSecAppTaskResponseInfoCode{
			value: "success",
		},
		FAILURE: DeleteSecAppTaskResponseInfoCode{
			value: "failure",
		},
	}
}

func (c DeleteSecAppTaskResponseInfoCode) Value() string {
	return c.value
}

func (c DeleteSecAppTaskResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSecAppTaskResponseInfoCode) UnmarshalJSON(b []byte) error {
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
