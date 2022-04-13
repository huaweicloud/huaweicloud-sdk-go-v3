package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type UpdateFalsePositiveResponse struct {
	// 状态码:   * success - 成功   * failure - 失败

	InfoCode *UpdateFalsePositiveResponseInfoCode `json:"info_code,omitempty"`
	// 返回的提示信息

	InfoDescription *string `json:"info_description,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o UpdateFalsePositiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFalsePositiveResponse struct{}"
	}

	return strings.Join([]string{"UpdateFalsePositiveResponse", string(data)}, " ")
}

type UpdateFalsePositiveResponseInfoCode struct {
	value string
}

type UpdateFalsePositiveResponseInfoCodeEnum struct {
	SUCCESS UpdateFalsePositiveResponseInfoCode
	FAILURE UpdateFalsePositiveResponseInfoCode
}

func GetUpdateFalsePositiveResponseInfoCodeEnum() UpdateFalsePositiveResponseInfoCodeEnum {
	return UpdateFalsePositiveResponseInfoCodeEnum{
		SUCCESS: UpdateFalsePositiveResponseInfoCode{
			value: "success",
		},
		FAILURE: UpdateFalsePositiveResponseInfoCode{
			value: "failure",
		},
	}
}

func (c UpdateFalsePositiveResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFalsePositiveResponseInfoCode) UnmarshalJSON(b []byte) error {
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
