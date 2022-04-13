package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

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

func (c OperateInfoResponseBodyInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperateInfoResponseBodyInfoCode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
