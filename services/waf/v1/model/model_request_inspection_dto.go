package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RequestInspectionDto 请求的检查规则
type RequestInspectionDto struct {

	// 请求的 Payload 类型（JSON/FORM_URLENCODED）
	PayloadType RequestInspectionDtoPayloadType `json:"payload_type"`

	// 用户名字段路径
	Username string `json:"username"`

	// 密码字段路径
	Password string `json:"password"`

	// 邮箱字段路径（可选）
	Email *string `json:"email,omitempty"`

	// 手机号字段路径（可选）
	Phone *string `json:"phone,omitempty"`
}

func (o RequestInspectionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestInspectionDto struct{}"
	}

	return strings.Join([]string{"RequestInspectionDto", string(data)}, " ")
}

type RequestInspectionDtoPayloadType struct {
	value string
}

type RequestInspectionDtoPayloadTypeEnum struct {
	JSON            RequestInspectionDtoPayloadType
	FORM_URLENCODED RequestInspectionDtoPayloadType
}

func GetRequestInspectionDtoPayloadTypeEnum() RequestInspectionDtoPayloadTypeEnum {
	return RequestInspectionDtoPayloadTypeEnum{
		JSON: RequestInspectionDtoPayloadType{
			value: "JSON",
		},
		FORM_URLENCODED: RequestInspectionDtoPayloadType{
			value: "FORM_URLENCODED",
		},
	}
}

func (c RequestInspectionDtoPayloadType) Value() string {
	return c.value
}

func (c RequestInspectionDtoPayloadType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RequestInspectionDtoPayloadType) UnmarshalJSON(b []byte) error {
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
