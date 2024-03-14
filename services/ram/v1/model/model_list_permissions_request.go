package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPermissionsRequest Request Object
type ListPermissionsRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 资源类型的名称。
	ResourceType *string `json:"resource_type,omitempty"`

	// 权限类型。RAM_MANAGED表示RAM托管的权限，CUSTOMER_MANAGED表示租户创建的自定义的权限，ALL表示所有的权限。
	PermissionType *ListPermissionsRequestPermissionType `json:"permission_type,omitempty"`
}

func (o ListPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionsRequest", string(data)}, " ")
}

type ListPermissionsRequestPermissionType struct {
	value string
}

type ListPermissionsRequestPermissionTypeEnum struct {
	RAM_MANAGED      ListPermissionsRequestPermissionType
	CUSTOMER_MANAGED ListPermissionsRequestPermissionType
	ALL              ListPermissionsRequestPermissionType
}

func GetListPermissionsRequestPermissionTypeEnum() ListPermissionsRequestPermissionTypeEnum {
	return ListPermissionsRequestPermissionTypeEnum{
		RAM_MANAGED: ListPermissionsRequestPermissionType{
			value: "RAM_MANAGED",
		},
		CUSTOMER_MANAGED: ListPermissionsRequestPermissionType{
			value: "CUSTOMER_MANAGED",
		},
		ALL: ListPermissionsRequestPermissionType{
			value: "ALL",
		},
	}
}

func (c ListPermissionsRequestPermissionType) Value() string {
	return c.value
}

func (c ListPermissionsRequestPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPermissionsRequestPermissionType) UnmarshalJSON(b []byte) error {
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
