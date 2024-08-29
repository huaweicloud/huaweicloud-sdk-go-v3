package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteAgencyWithRoleTypeRequest Request Object
type DeleteAgencyWithRoleTypeRequest struct {

	// 委托授权类型 * CBS:对话机器人服务(CBS)访客 * SIS:语音交互服务(SIS)调用
	RoleType DeleteAgencyWithRoleTypeRequestRoleType `json:"role_type"`
}

func (o DeleteAgencyWithRoleTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyWithRoleTypeRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgencyWithRoleTypeRequest", string(data)}, " ")
}

type DeleteAgencyWithRoleTypeRequestRoleType struct {
	value string
}

type DeleteAgencyWithRoleTypeRequestRoleTypeEnum struct {
	CBS DeleteAgencyWithRoleTypeRequestRoleType
	SIS DeleteAgencyWithRoleTypeRequestRoleType
}

func GetDeleteAgencyWithRoleTypeRequestRoleTypeEnum() DeleteAgencyWithRoleTypeRequestRoleTypeEnum {
	return DeleteAgencyWithRoleTypeRequestRoleTypeEnum{
		CBS: DeleteAgencyWithRoleTypeRequestRoleType{
			value: "CBS",
		},
		SIS: DeleteAgencyWithRoleTypeRequestRoleType{
			value: "SIS",
		},
	}
}

func (c DeleteAgencyWithRoleTypeRequestRoleType) Value() string {
	return c.value
}

func (c DeleteAgencyWithRoleTypeRequestRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteAgencyWithRoleTypeRequestRoleType) UnmarshalJSON(b []byte) error {
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
