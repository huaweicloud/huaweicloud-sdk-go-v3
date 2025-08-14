package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplicationDto 应用程序
type ApplicationDto struct {

	// 应用程序URN
	ApplicationUrn *string `json:"application_urn,omitempty"`

	// 应用程序提供商URN
	ApplicationProviderUrn *string `json:"application_provider_urn,omitempty"`

	AssignmentConfig *AssignmentConfigDto `json:"assignment_config,omitempty"`

	// 应用程序创建时间
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 应用程序描述
	Description *string `json:"description,omitempty"`

	// IAM身份中心实例URN
	InstanceUrn *string `json:"instance_urn,omitempty"`

	// 应用程序显示名称
	Name *string `json:"name,omitempty"`

	PortalOptions *PortalOptionsDto `json:"portal_options,omitempty"`

	// 应用程序状态
	Status *ApplicationDtoStatus `json:"status,omitempty"`

	// 华为云账号
	ApplicationAccount *string `json:"application_account,omitempty"`
}

func (o ApplicationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationDto struct{}"
	}

	return strings.Join([]string{"ApplicationDto", string(data)}, " ")
}

type ApplicationDtoStatus struct {
	value string
}

type ApplicationDtoStatusEnum struct {
	DISABLED ApplicationDtoStatus
	ENABLED  ApplicationDtoStatus
}

func GetApplicationDtoStatusEnum() ApplicationDtoStatusEnum {
	return ApplicationDtoStatusEnum{
		DISABLED: ApplicationDtoStatus{
			value: "DISABLED",
		},
		ENABLED: ApplicationDtoStatus{
			value: "ENABLED",
		},
	}
}

func (c ApplicationDtoStatus) Value() string {
	return c.value
}

func (c ApplicationDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationDtoStatus) UnmarshalJSON(b []byte) error {
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
