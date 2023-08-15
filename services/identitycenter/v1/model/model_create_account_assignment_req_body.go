package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAccountAssignmentReqBody CreateAccountAssignment请求体
type CreateAccountAssignmentReqBody struct {

	// 权限集唯一标识
	PermissionSetId string `json:"permission_set_id"`

	// IAM身份中心中的一个实体身份唯一标识，例如用户或用户组
	PrincipalId string `json:"principal_id"`

	// 创建绑定的实体类型
	PrincipalType CreateAccountAssignmentReqBodyPrincipalType `json:"principal_type"`

	// 待绑定的目标实体标识.
	TargetId string `json:"target_id"`

	// 创建绑定的实体类型
	TargetType CreateAccountAssignmentReqBodyTargetType `json:"target_type"`
}

func (o CreateAccountAssignmentReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountAssignmentReqBody struct{}"
	}

	return strings.Join([]string{"CreateAccountAssignmentReqBody", string(data)}, " ")
}

type CreateAccountAssignmentReqBodyPrincipalType struct {
	value string
}

type CreateAccountAssignmentReqBodyPrincipalTypeEnum struct {
	USER  CreateAccountAssignmentReqBodyPrincipalType
	GROUP CreateAccountAssignmentReqBodyPrincipalType
}

func GetCreateAccountAssignmentReqBodyPrincipalTypeEnum() CreateAccountAssignmentReqBodyPrincipalTypeEnum {
	return CreateAccountAssignmentReqBodyPrincipalTypeEnum{
		USER: CreateAccountAssignmentReqBodyPrincipalType{
			value: "USER",
		},
		GROUP: CreateAccountAssignmentReqBodyPrincipalType{
			value: "GROUP",
		},
	}
}

func (c CreateAccountAssignmentReqBodyPrincipalType) Value() string {
	return c.value
}

func (c CreateAccountAssignmentReqBodyPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccountAssignmentReqBodyPrincipalType) UnmarshalJSON(b []byte) error {
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

type CreateAccountAssignmentReqBodyTargetType struct {
	value string
}

type CreateAccountAssignmentReqBodyTargetTypeEnum struct {
	ACCOUNT CreateAccountAssignmentReqBodyTargetType
}

func GetCreateAccountAssignmentReqBodyTargetTypeEnum() CreateAccountAssignmentReqBodyTargetTypeEnum {
	return CreateAccountAssignmentReqBodyTargetTypeEnum{
		ACCOUNT: CreateAccountAssignmentReqBodyTargetType{
			value: "ACCOUNT",
		},
	}
}

func (c CreateAccountAssignmentReqBodyTargetType) Value() string {
	return c.value
}

func (c CreateAccountAssignmentReqBodyTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccountAssignmentReqBodyTargetType) UnmarshalJSON(b []byte) error {
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
