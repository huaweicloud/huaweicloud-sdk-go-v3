package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ValidationLinkConfig struct {

	// 错误描述
	Message *string `json:"message,omitempty" xml:"message"`

	// ERROR,WARNING
	Status *ValidationLinkConfigStatus `json:"status,omitempty" xml:"status"`
}

func (o ValidationLinkConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidationLinkConfig struct{}"
	}

	return strings.Join([]string{"ValidationLinkConfig", string(data)}, " ")
}

type ValidationLinkConfigStatus struct {
	value string
}

type ValidationLinkConfigStatusEnum struct {
	ERROR   ValidationLinkConfigStatus
	WARNING ValidationLinkConfigStatus
}

func GetValidationLinkConfigStatusEnum() ValidationLinkConfigStatusEnum {
	return ValidationLinkConfigStatusEnum{
		ERROR: ValidationLinkConfigStatus{
			value: "ERROR",
		},
		WARNING: ValidationLinkConfigStatus{
			value: "WARNING",
		},
	}
}

func (c ValidationLinkConfigStatus) Value() string {
	return c.value
}

func (c ValidationLinkConfigStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidationLinkConfigStatus) UnmarshalJSON(b []byte) error {
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
