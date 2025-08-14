package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateApplicationAssignmentReqBody CreateApplicationAssignment的请求体
type CreateApplicationAssignmentReqBody struct {

	// 主体Id
	PrincipalId string `json:"principal_id"`

	// 主体类型
	PrincipalType CreateApplicationAssignmentReqBodyPrincipalType `json:"principal_type"`
}

func (o CreateApplicationAssignmentReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationAssignmentReqBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationAssignmentReqBody", string(data)}, " ")
}

type CreateApplicationAssignmentReqBodyPrincipalType struct {
	value string
}

type CreateApplicationAssignmentReqBodyPrincipalTypeEnum struct {
	USER  CreateApplicationAssignmentReqBodyPrincipalType
	GROUP CreateApplicationAssignmentReqBodyPrincipalType
}

func GetCreateApplicationAssignmentReqBodyPrincipalTypeEnum() CreateApplicationAssignmentReqBodyPrincipalTypeEnum {
	return CreateApplicationAssignmentReqBodyPrincipalTypeEnum{
		USER: CreateApplicationAssignmentReqBodyPrincipalType{
			value: "USER",
		},
		GROUP: CreateApplicationAssignmentReqBodyPrincipalType{
			value: "GROUP",
		},
	}
}

func (c CreateApplicationAssignmentReqBodyPrincipalType) Value() string {
	return c.value
}

func (c CreateApplicationAssignmentReqBodyPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApplicationAssignmentReqBodyPrincipalType) UnmarshalJSON(b []byte) error {
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
