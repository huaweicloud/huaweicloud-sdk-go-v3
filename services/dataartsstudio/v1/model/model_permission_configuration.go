package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionConfiguration struct {

	// 数据源类型
	DatasourceType *string `json:"datasource_type,omitempty"`

	// 数据源操作权限类型
	PermissionTypes *[]PermissionConfigurationPermissionTypes `json:"permission_types,omitempty"`

	// 数据源支持的权限操作列表
	PermissionActions *[]PermissionActions `json:"permission_actions,omitempty"`
}

func (o PermissionConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionConfiguration struct{}"
	}

	return strings.Join([]string{"PermissionConfiguration", string(data)}, " ")
}

type PermissionConfigurationPermissionTypes struct {
	value string
}

type PermissionConfigurationPermissionTypesEnum struct {
	ALLOW PermissionConfigurationPermissionTypes
}

func GetPermissionConfigurationPermissionTypesEnum() PermissionConfigurationPermissionTypesEnum {
	return PermissionConfigurationPermissionTypesEnum{
		ALLOW: PermissionConfigurationPermissionTypes{
			value: "ALLOW",
		},
	}
}

func (c PermissionConfigurationPermissionTypes) Value() string {
	return c.value
}

func (c PermissionConfigurationPermissionTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionConfigurationPermissionTypes) UnmarshalJSON(b []byte) error {
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
