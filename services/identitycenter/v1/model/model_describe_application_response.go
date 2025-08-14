package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DescribeApplicationResponse Response Object
type DescribeApplicationResponse struct {

	// 应用程序URN
	ApplicationUrn *string `json:"application_urn,omitempty"`

	// 应用程序提供商URN
	ApplicationProviderUrn *string `json:"application_provider_urn,omitempty"`

	AssignmentConfig *AssignmentConfigDto `json:"assignment_config,omitempty"`

	// 创建时间
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 应用程序描述
	Description *string `json:"description,omitempty"`

	// IAM身份中心实例URN
	InstanceUrn *string `json:"instance_urn,omitempty"`

	// 应用程序显示名
	Name *string `json:"name,omitempty"`

	PortalOptions *PortalOptionsDto `json:"portal_options,omitempty"`

	// 状态
	Status *DescribeApplicationResponseStatus `json:"status,omitempty"`

	// 华为云账号
	ApplicationAccount *string `json:"application_account,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o DescribeApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeApplicationResponse struct{}"
	}

	return strings.Join([]string{"DescribeApplicationResponse", string(data)}, " ")
}

type DescribeApplicationResponseStatus struct {
	value string
}

type DescribeApplicationResponseStatusEnum struct {
	ENABLED  DescribeApplicationResponseStatus
	DISABLED DescribeApplicationResponseStatus
}

func GetDescribeApplicationResponseStatusEnum() DescribeApplicationResponseStatusEnum {
	return DescribeApplicationResponseStatusEnum{
		ENABLED: DescribeApplicationResponseStatus{
			value: "ENABLED",
		},
		DISABLED: DescribeApplicationResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c DescribeApplicationResponseStatus) Value() string {
	return c.value
}

func (c DescribeApplicationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DescribeApplicationResponseStatus) UnmarshalJSON(b []byte) error {
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
