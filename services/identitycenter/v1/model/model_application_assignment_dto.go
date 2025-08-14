package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplicationAssignmentDto 应用程序分配的主体
type ApplicationAssignmentDto struct {

	// 应用程序URN
	ApplicationUrn string `json:"application_urn"`

	// 主体Id
	PrincipalId string `json:"principal_id"`

	// 主体类型
	PrincipalType ApplicationAssignmentDtoPrincipalType `json:"principal_type"`
}

func (o ApplicationAssignmentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationAssignmentDto struct{}"
	}

	return strings.Join([]string{"ApplicationAssignmentDto", string(data)}, " ")
}

type ApplicationAssignmentDtoPrincipalType struct {
	value string
}

type ApplicationAssignmentDtoPrincipalTypeEnum struct {
	USER  ApplicationAssignmentDtoPrincipalType
	GROUP ApplicationAssignmentDtoPrincipalType
}

func GetApplicationAssignmentDtoPrincipalTypeEnum() ApplicationAssignmentDtoPrincipalTypeEnum {
	return ApplicationAssignmentDtoPrincipalTypeEnum{
		USER: ApplicationAssignmentDtoPrincipalType{
			value: "USER",
		},
		GROUP: ApplicationAssignmentDtoPrincipalType{
			value: "GROUP",
		},
	}
}

func (c ApplicationAssignmentDtoPrincipalType) Value() string {
	return c.value
}

func (c ApplicationAssignmentDtoPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationAssignmentDtoPrincipalType) UnmarshalJSON(b []byte) error {
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
