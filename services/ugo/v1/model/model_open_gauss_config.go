package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussConfig GaussDB的特殊配置。
type OpenGaussConfig struct {

	// 权限检查类型。
	PermissionCheckType *OpenGaussConfigPermissionCheckType `json:"permission_check_type,omitempty"`
}

func (o OpenGaussConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussConfig struct{}"
	}

	return strings.Join([]string{"OpenGaussConfig", string(data)}, " ")
}

type OpenGaussConfigPermissionCheckType struct {
	value string
}

type OpenGaussConfigPermissionCheckTypeEnum struct {
	OBJECTOWNER OpenGaussConfigPermissionCheckType
	SYSADMIN    OpenGaussConfigPermissionCheckType
}

func GetOpenGaussConfigPermissionCheckTypeEnum() OpenGaussConfigPermissionCheckTypeEnum {
	return OpenGaussConfigPermissionCheckTypeEnum{
		OBJECTOWNER: OpenGaussConfigPermissionCheckType{
			value: "objectowner",
		},
		SYSADMIN: OpenGaussConfigPermissionCheckType{
			value: "sysadmin",
		},
	}
}

func (c OpenGaussConfigPermissionCheckType) Value() string {
	return c.value
}

func (c OpenGaussConfigPermissionCheckType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussConfigPermissionCheckType) UnmarshalJSON(b []byte) error {
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
