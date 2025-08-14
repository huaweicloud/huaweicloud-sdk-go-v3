package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PortalOptionsDto 门户选项设置
type PortalOptionsDto struct {

	// 应用程序在门户是否可见
	Visible *bool `json:"visible,omitempty"`

	// 应用程序在门户可见性
	Visibility *PortalOptionsDtoVisibility `json:"visibility,omitempty"`

	SignInOptions *SignInOptionsDto `json:"sign_in_options,omitempty"`
}

func (o PortalOptionsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortalOptionsDto struct{}"
	}

	return strings.Join([]string{"PortalOptionsDto", string(data)}, " ")
}

type PortalOptionsDtoVisibility struct {
	value string
}

type PortalOptionsDtoVisibilityEnum struct {
	ENABLED  PortalOptionsDtoVisibility
	DISABLED PortalOptionsDtoVisibility
}

func GetPortalOptionsDtoVisibilityEnum() PortalOptionsDtoVisibilityEnum {
	return PortalOptionsDtoVisibilityEnum{
		ENABLED: PortalOptionsDtoVisibility{
			value: "ENABLED",
		},
		DISABLED: PortalOptionsDtoVisibility{
			value: "DISABLED",
		},
	}
}

func (c PortalOptionsDtoVisibility) Value() string {
	return c.value
}

func (c PortalOptionsDtoVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortalOptionsDtoVisibility) UnmarshalJSON(b []byte) error {
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
