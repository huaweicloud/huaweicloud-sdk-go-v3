package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAgencyWithRoleTypeRequest Request Object
type CreateAgencyWithRoleTypeRequest struct {

	// 委托授权类型 * CBS:对话机器人服务(CBS)访客 * SIS:语音交互服务(SIS)调用
	RoleType CreateAgencyWithRoleTypeRequestRoleType `json:"role_type"`
}

func (o CreateAgencyWithRoleTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyWithRoleTypeRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyWithRoleTypeRequest", string(data)}, " ")
}

type CreateAgencyWithRoleTypeRequestRoleType struct {
	value string
}

type CreateAgencyWithRoleTypeRequestRoleTypeEnum struct {
	CBS CreateAgencyWithRoleTypeRequestRoleType
	SIS CreateAgencyWithRoleTypeRequestRoleType
}

func GetCreateAgencyWithRoleTypeRequestRoleTypeEnum() CreateAgencyWithRoleTypeRequestRoleTypeEnum {
	return CreateAgencyWithRoleTypeRequestRoleTypeEnum{
		CBS: CreateAgencyWithRoleTypeRequestRoleType{
			value: "CBS",
		},
		SIS: CreateAgencyWithRoleTypeRequestRoleType{
			value: "SIS",
		},
	}
}

func (c CreateAgencyWithRoleTypeRequestRoleType) Value() string {
	return c.value
}

func (c CreateAgencyWithRoleTypeRequestRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAgencyWithRoleTypeRequestRoleType) UnmarshalJSON(b []byte) error {
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
