package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperateInfoResponseBody struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *OperateInfoResponseBodyInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`
}

func (o OperateInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateInfoResponseBody struct{}"
	}

	return strings.Join([]string{"OperateInfoResponseBody", string(data)}, " ")
}

type OperateInfoResponseBodyInfoCode struct {
	value string
}

type OperateInfoResponseBodyInfoCodeEnum struct {
	SUCCESS OperateInfoResponseBodyInfoCode
	FAILURE OperateInfoResponseBodyInfoCode
}

func GetOperateInfoResponseBodyInfoCodeEnum() OperateInfoResponseBodyInfoCodeEnum {
	return OperateInfoResponseBodyInfoCodeEnum{
		SUCCESS: OperateInfoResponseBodyInfoCode{
			value: "success",
		},
		FAILURE: OperateInfoResponseBodyInfoCode{
			value: "failure",
		},
	}
}

func (c OperateInfoResponseBodyInfoCode) Value() string {
	return c.value
}

func (c OperateInfoResponseBodyInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperateInfoResponseBodyInfoCode) UnmarshalJSON(b []byte) error {
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
