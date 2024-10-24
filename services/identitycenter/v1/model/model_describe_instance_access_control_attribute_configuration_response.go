package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DescribeInstanceAccessControlAttributeConfigurationResponse Response Object
type DescribeInstanceAccessControlAttributeConfigurationResponse struct {
	InstanceAccessControlAttributeConfiguration *InstanceAccessControlAttributeConfigurationDto `json:"instance_access_control_attribute_configuration,omitempty"`

	// ABAC属性配置的状态
	Status *DescribeInstanceAccessControlAttributeConfigurationResponseStatus `json:"status,omitempty"`

	// 提供有关指定属性的当前状态的更多详细信息
	StatusReason   *string `json:"status_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DescribeInstanceAccessControlAttributeConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeInstanceAccessControlAttributeConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DescribeInstanceAccessControlAttributeConfigurationResponse", string(data)}, " ")
}

type DescribeInstanceAccessControlAttributeConfigurationResponseStatus struct {
	value string
}

type DescribeInstanceAccessControlAttributeConfigurationResponseStatusEnum struct {
	ENABLED              DescribeInstanceAccessControlAttributeConfigurationResponseStatus
	CREATION_IN_PROGRESS DescribeInstanceAccessControlAttributeConfigurationResponseStatus
	CREATION_FAILED      DescribeInstanceAccessControlAttributeConfigurationResponseStatus
}

func GetDescribeInstanceAccessControlAttributeConfigurationResponseStatusEnum() DescribeInstanceAccessControlAttributeConfigurationResponseStatusEnum {
	return DescribeInstanceAccessControlAttributeConfigurationResponseStatusEnum{
		ENABLED: DescribeInstanceAccessControlAttributeConfigurationResponseStatus{
			value: "ENABLED",
		},
		CREATION_IN_PROGRESS: DescribeInstanceAccessControlAttributeConfigurationResponseStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: DescribeInstanceAccessControlAttributeConfigurationResponseStatus{
			value: "CREATION_FAILED",
		},
	}
}

func (c DescribeInstanceAccessControlAttributeConfigurationResponseStatus) Value() string {
	return c.value
}

func (c DescribeInstanceAccessControlAttributeConfigurationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DescribeInstanceAccessControlAttributeConfigurationResponseStatus) UnmarshalJSON(b []byte) error {
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
