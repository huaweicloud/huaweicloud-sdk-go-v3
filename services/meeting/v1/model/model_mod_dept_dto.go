package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ModDeptDto struct {
	// 部门名称 maxLength：128 minLength：1

	DeptName *string `json:"deptName,omitempty"`
	// 父部门编码 maxLength：32

	ParentDeptCode *string `json:"parentDeptCode,omitempty"`
	// 备注 maxLength：96 minLength：0

	Note *string `json:"note,omitempty"`
	// 其他用户对该部门下用户的访问权限： - UNLIMITED：默认，不做限制 - OPEN：公开，其他部门都可访问（无论对方权限如何配置） - CLOSE：隐藏，其他部门不可访问（暂未实现） - DESIGNATED_DEPARTMENT：指定部门能访问（暂未实现）

	InPermission *ModDeptDtoInPermission `json:"inPermission,omitempty"`
	// 该部门下用户访问权限控制 - UNLIMITED：不限制 - ONLY_SELF：仅能查询自己 - SELF_AND_CHILD_DEPARTMENT：该部门下用户能查询本部门及子部门通讯 - DESIGNATED_DEPARTMENT：该部门下用户能查询指定部门通讯录

	OutPermission *ModDeptDtoOutPermission `json:"outPermission,omitempty"`
	// 允许访问的部门列表,仅outPermission为DESIGNATED_DEPARTMENT时有效，最多支持配置150

	DesignatedOutDeptCodes *[]string `json:"designatedOutDeptCodes,omitempty"`
}

func (o ModDeptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModDeptDto struct{}"
	}

	return strings.Join([]string{"ModDeptDto", string(data)}, " ")
}

type ModDeptDtoInPermission struct {
	value string
}

type ModDeptDtoInPermissionEnum struct {
	UNLIMITED             ModDeptDtoInPermission
	OPEN                  ModDeptDtoInPermission
	CLOSE                 ModDeptDtoInPermission
	DESIGNATED_DEPARTMENT ModDeptDtoInPermission
}

func GetModDeptDtoInPermissionEnum() ModDeptDtoInPermissionEnum {
	return ModDeptDtoInPermissionEnum{
		UNLIMITED: ModDeptDtoInPermission{
			value: "UNLIMITED",
		},
		OPEN: ModDeptDtoInPermission{
			value: "OPEN",
		},
		CLOSE: ModDeptDtoInPermission{
			value: "CLOSE",
		},
		DESIGNATED_DEPARTMENT: ModDeptDtoInPermission{
			value: "DESIGNATED_DEPARTMENT",
		},
	}
}

func (c ModDeptDtoInPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModDeptDtoInPermission) UnmarshalJSON(b []byte) error {
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

type ModDeptDtoOutPermission struct {
	value string
}

type ModDeptDtoOutPermissionEnum struct {
	UNLIMITED                 ModDeptDtoOutPermission
	ONLY_SELF                 ModDeptDtoOutPermission
	SELF_AND_CHILD_DEPARTMENT ModDeptDtoOutPermission
	DESIGNATED_DEPARTMENT     ModDeptDtoOutPermission
}

func GetModDeptDtoOutPermissionEnum() ModDeptDtoOutPermissionEnum {
	return ModDeptDtoOutPermissionEnum{
		UNLIMITED: ModDeptDtoOutPermission{
			value: "UNLIMITED",
		},
		ONLY_SELF: ModDeptDtoOutPermission{
			value: "ONLY_SELF",
		},
		SELF_AND_CHILD_DEPARTMENT: ModDeptDtoOutPermission{
			value: "SELF_AND_CHILD_DEPARTMENT",
		},
		DESIGNATED_DEPARTMENT: ModDeptDtoOutPermission{
			value: "DESIGNATED_DEPARTMENT",
		},
	}
}

func (c ModDeptDtoOutPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModDeptDtoOutPermission) UnmarshalJSON(b []byte) error {
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
