package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteApplicationAssignmentReqBody DeleteApplicationAssignment的请求体
type DeleteApplicationAssignmentReqBody struct {

	// 主体Id
	PrincipalId string `json:"principal_id"`

	// 主体类型
	PrincipalType DeleteApplicationAssignmentReqBodyPrincipalType `json:"principal_type"`
}

func (o DeleteApplicationAssignmentReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationAssignmentReqBody struct{}"
	}

	return strings.Join([]string{"DeleteApplicationAssignmentReqBody", string(data)}, " ")
}

type DeleteApplicationAssignmentReqBodyPrincipalType struct {
	value string
}

type DeleteApplicationAssignmentReqBodyPrincipalTypeEnum struct {
	USER  DeleteApplicationAssignmentReqBodyPrincipalType
	GROUP DeleteApplicationAssignmentReqBodyPrincipalType
}

func GetDeleteApplicationAssignmentReqBodyPrincipalTypeEnum() DeleteApplicationAssignmentReqBodyPrincipalTypeEnum {
	return DeleteApplicationAssignmentReqBodyPrincipalTypeEnum{
		USER: DeleteApplicationAssignmentReqBodyPrincipalType{
			value: "USER",
		},
		GROUP: DeleteApplicationAssignmentReqBodyPrincipalType{
			value: "GROUP",
		},
	}
}

func (c DeleteApplicationAssignmentReqBodyPrincipalType) Value() string {
	return c.value
}

func (c DeleteApplicationAssignmentReqBodyPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteApplicationAssignmentReqBodyPrincipalType) UnmarshalJSON(b []byte) error {
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
