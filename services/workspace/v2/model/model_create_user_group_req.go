package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateUserGroupReq struct {

	// 用户组名称。
	GroupName string `json:"group_name"`

	// 用户组类型。 * AD： AD域用户组 * LOCAL： 本地liteAs用户组
	PlatformType CreateUserGroupReqPlatformType `json:"platform_type"`

	// 用户组描述。
	Description *string `json:"description,omitempty"`
}

func (o CreateUserGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserGroupReq struct{}"
	}

	return strings.Join([]string{"CreateUserGroupReq", string(data)}, " ")
}

type CreateUserGroupReqPlatformType struct {
	value string
}

type CreateUserGroupReqPlatformTypeEnum struct {
	AD    CreateUserGroupReqPlatformType
	LOCAL CreateUserGroupReqPlatformType
}

func GetCreateUserGroupReqPlatformTypeEnum() CreateUserGroupReqPlatformTypeEnum {
	return CreateUserGroupReqPlatformTypeEnum{
		AD: CreateUserGroupReqPlatformType{
			value: "AD",
		},
		LOCAL: CreateUserGroupReqPlatformType{
			value: "LOCAL",
		},
	}
}

func (c CreateUserGroupReqPlatformType) Value() string {
	return c.value
}

func (c CreateUserGroupReqPlatformType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserGroupReqPlatformType) UnmarshalJSON(b []byte) error {
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
