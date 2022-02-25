package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeptDto struct {
	// 部门编码，企业内唯一，若携带则以携带为准，不支持修改。 maxLength：32

	DeptCode *string `json:"deptCode,omitempty"`
	// 部门名称 maxLength：128 minLength：1

	DeptName string `json:"deptName"`
	// 父部门编码,默认为根部门。 默认值：1： maxLength：32

	ParentDeptCode *string `json:"parentDeptCode,omitempty"`
	// 备注 maxLength：96 minLength：0

	Note *string `json:"note,omitempty"`
	// 其他用户对该部门下用户的访问权限： - UNLIMITED：默认，不做限制 - OPEN：公开，其他部门都可访问（无论对方权限如何配置）

	InPermission *DeptDtoInPermission `json:"inPermission,omitempty"`
	// 该部门下用户访问权限控制 - UNLIMITED：不限制 - ONLY_SELF：仅能查询自己 - SELF_AND_CHILD_DEPARTMENT：该部门下用户能查询本部门及子部门通讯 - DESIGNATED_DEPARTMENT：该部门下用户能查询指定部门通讯录

	OutPermission *DeptDtoOutPermission `json:"outPermission,omitempty"`
	// 允许访问的部门列表,仅outPermission为DESIGNATED_DEPARTMENT时有效，最多支持配置150

	DesignatedOutDeptCodes *[]string `json:"designatedOutDeptCodes,omitempty"`
}

func (o DeptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeptDto struct{}"
	}

	return strings.Join([]string{"DeptDto", string(data)}, " ")
}

type DeptDtoInPermission struct {
	value string
}

type DeptDtoInPermissionEnum struct {
	UNLIMITED             DeptDtoInPermission
	OPEN                  DeptDtoInPermission
	CLOSE                 DeptDtoInPermission
	DESIGNATED_DEPARTMENT DeptDtoInPermission
}

func GetDeptDtoInPermissionEnum() DeptDtoInPermissionEnum {
	return DeptDtoInPermissionEnum{
		UNLIMITED: DeptDtoInPermission{
			value: "UNLIMITED",
		},
		OPEN: DeptDtoInPermission{
			value: "OPEN",
		},
		CLOSE: DeptDtoInPermission{
			value: "CLOSE",
		},
		DESIGNATED_DEPARTMENT: DeptDtoInPermission{
			value: "DESIGNATED_DEPARTMENT",
		},
	}
}

func (c DeptDtoInPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeptDtoInPermission) UnmarshalJSON(b []byte) error {
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

type DeptDtoOutPermission struct {
	value string
}

type DeptDtoOutPermissionEnum struct {
	UNLIMITED                 DeptDtoOutPermission
	ONLY_SELF                 DeptDtoOutPermission
	SELF_AND_CHILD_DEPARTMENT DeptDtoOutPermission
	DESIGNATED_DEPARTMENT     DeptDtoOutPermission
}

func GetDeptDtoOutPermissionEnum() DeptDtoOutPermissionEnum {
	return DeptDtoOutPermissionEnum{
		UNLIMITED: DeptDtoOutPermission{
			value: "UNLIMITED",
		},
		ONLY_SELF: DeptDtoOutPermission{
			value: "ONLY_SELF",
		},
		SELF_AND_CHILD_DEPARTMENT: DeptDtoOutPermission{
			value: "SELF_AND_CHILD_DEPARTMENT",
		},
		DESIGNATED_DEPARTMENT: DeptDtoOutPermission{
			value: "DESIGNATED_DEPARTMENT",
		},
	}
}

func (c DeptDtoOutPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeptDtoOutPermission) UnmarshalJSON(b []byte) error {
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
