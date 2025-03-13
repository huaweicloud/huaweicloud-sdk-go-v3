package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionApprovalDetailDtoPermissions struct {

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 权限
	PermissionAction *[]PermissionApprovalDetailDtoPermissionsPermissionAction `json:"permission_action,omitempty"`

	// 申请权限所在的空间权限集
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 密级
	SecrecyLevelId *string `json:"secrecy_level_id,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`
}

func (o PermissionApprovalDetailDtoPermissions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionApprovalDetailDtoPermissions struct{}"
	}

	return strings.Join([]string{"PermissionApprovalDetailDtoPermissions", string(data)}, " ")
}

type PermissionApprovalDetailDtoPermissionsPermissionAction struct {
	value string
}

type PermissionApprovalDetailDtoPermissionsPermissionActionEnum struct {
	SELECT PermissionApprovalDetailDtoPermissionsPermissionAction
}

func GetPermissionApprovalDetailDtoPermissionsPermissionActionEnum() PermissionApprovalDetailDtoPermissionsPermissionActionEnum {
	return PermissionApprovalDetailDtoPermissionsPermissionActionEnum{
		SELECT: PermissionApprovalDetailDtoPermissionsPermissionAction{
			value: "SELECT",
		},
	}
}

func (c PermissionApprovalDetailDtoPermissionsPermissionAction) Value() string {
	return c.value
}

func (c PermissionApprovalDetailDtoPermissionsPermissionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionApprovalDetailDtoPermissionsPermissionAction) UnmarshalJSON(b []byte) error {
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
