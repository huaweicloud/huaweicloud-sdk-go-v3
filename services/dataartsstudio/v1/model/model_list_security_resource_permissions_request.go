package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityResourcePermissionsRequest Request Object
type ListSecurityResourcePermissionsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 策略名称。
	PolicyName *string `json:"policy_name,omitempty"`

	// 授权资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 成员名称
	MemberName *string `json:"member_name,omitempty"`

	// 排序参数, NAME,UPDATE_TIME。
	OrderBy *ListSecurityResourcePermissionsRequestOrderBy `json:"order_by,omitempty"`

	// 是否升序（仅指定排序参数时有效）
	OrderByAsc *bool `json:"order_by_asc,omitempty"`
}

func (o ListSecurityResourcePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityResourcePermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityResourcePermissionsRequest", string(data)}, " ")
}

type ListSecurityResourcePermissionsRequestOrderBy struct {
	value string
}

type ListSecurityResourcePermissionsRequestOrderByEnum struct {
	NAME        ListSecurityResourcePermissionsRequestOrderBy
	UPDATE_TIME ListSecurityResourcePermissionsRequestOrderBy
}

func GetListSecurityResourcePermissionsRequestOrderByEnum() ListSecurityResourcePermissionsRequestOrderByEnum {
	return ListSecurityResourcePermissionsRequestOrderByEnum{
		NAME: ListSecurityResourcePermissionsRequestOrderBy{
			value: "NAME",
		},
		UPDATE_TIME: ListSecurityResourcePermissionsRequestOrderBy{
			value: "UPDATE_TIME",
		},
	}
}

func (c ListSecurityResourcePermissionsRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityResourcePermissionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityResourcePermissionsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
