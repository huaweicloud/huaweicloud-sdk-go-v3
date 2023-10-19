package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HostedCloudEnum - HWCloud (华为云) - Ireland (爱尔兰)
type HostedCloudEnum struct {
	value string
}

type HostedCloudEnumEnum struct {
	HW_CLOUD HostedCloudEnum
	IRELAND  HostedCloudEnum
}

func GetHostedCloudEnumEnum() HostedCloudEnumEnum {
	return HostedCloudEnumEnum{
		HW_CLOUD: HostedCloudEnum{
			value: "HWCloud",
		},
		IRELAND: HostedCloudEnum{
			value: "Ireland",
		},
	}
}

func (c HostedCloudEnum) Value() string {
	return c.value
}

func (c HostedCloudEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedCloudEnum) UnmarshalJSON(b []byte) error {
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
