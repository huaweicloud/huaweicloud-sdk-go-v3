package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteAccountAssignmentReqBody 解绑关联关系请求体
type DeleteAccountAssignmentReqBody struct {

	// 权限集唯一标识
	PermissionSetId string `json:"permission_set_id"`

	// IAM身份中心中的一个实体身份唯一标识，例如用户或用户组
	PrincipalId string `json:"principal_id"`

	// 实体类型
	PrincipalType DeleteAccountAssignmentReqBodyPrincipalType `json:"principal_type"`

	// 目标账户身份标识
	TargetId string `json:"target_id"`

	// 目标类型
	TargetType DeleteAccountAssignmentReqBodyTargetType `json:"target_type"`
}

func (o DeleteAccountAssignmentReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccountAssignmentReqBody struct{}"
	}

	return strings.Join([]string{"DeleteAccountAssignmentReqBody", string(data)}, " ")
}

type DeleteAccountAssignmentReqBodyPrincipalType struct {
	value string
}

type DeleteAccountAssignmentReqBodyPrincipalTypeEnum struct {
	USER  DeleteAccountAssignmentReqBodyPrincipalType
	GROUP DeleteAccountAssignmentReqBodyPrincipalType
}

func GetDeleteAccountAssignmentReqBodyPrincipalTypeEnum() DeleteAccountAssignmentReqBodyPrincipalTypeEnum {
	return DeleteAccountAssignmentReqBodyPrincipalTypeEnum{
		USER: DeleteAccountAssignmentReqBodyPrincipalType{
			value: "USER",
		},
		GROUP: DeleteAccountAssignmentReqBodyPrincipalType{
			value: "GROUP",
		},
	}
}

func (c DeleteAccountAssignmentReqBodyPrincipalType) Value() string {
	return c.value
}

func (c DeleteAccountAssignmentReqBodyPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteAccountAssignmentReqBodyPrincipalType) UnmarshalJSON(b []byte) error {
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

type DeleteAccountAssignmentReqBodyTargetType struct {
	value string
}

type DeleteAccountAssignmentReqBodyTargetTypeEnum struct {
	ACCOUNT DeleteAccountAssignmentReqBodyTargetType
}

func GetDeleteAccountAssignmentReqBodyTargetTypeEnum() DeleteAccountAssignmentReqBodyTargetTypeEnum {
	return DeleteAccountAssignmentReqBodyTargetTypeEnum{
		ACCOUNT: DeleteAccountAssignmentReqBodyTargetType{
			value: "ACCOUNT",
		},
	}
}

func (c DeleteAccountAssignmentReqBodyTargetType) Value() string {
	return c.value
}

func (c DeleteAccountAssignmentReqBodyTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteAccountAssignmentReqBodyTargetType) UnmarshalJSON(b []byte) error {
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
