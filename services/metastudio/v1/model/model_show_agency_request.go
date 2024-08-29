package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAgencyRequest Request Object
type ShowAgencyRequest struct {

	// 委托授权类型 * CBS:对话机器人服务(CBS)访客 * SIS:语音交互服务(SIS)调用
	RoleType *ShowAgencyRequestRoleType `json:"role_type,omitempty"`
}

func (o ShowAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyRequest", string(data)}, " ")
}

type ShowAgencyRequestRoleType struct {
	value string
}

type ShowAgencyRequestRoleTypeEnum struct {
	CBS ShowAgencyRequestRoleType
	SIS ShowAgencyRequestRoleType
}

func GetShowAgencyRequestRoleTypeEnum() ShowAgencyRequestRoleTypeEnum {
	return ShowAgencyRequestRoleTypeEnum{
		CBS: ShowAgencyRequestRoleType{
			value: "CBS",
		},
		SIS: ShowAgencyRequestRoleType{
			value: "SIS",
		},
	}
}

func (c ShowAgencyRequestRoleType) Value() string {
	return c.value
}

func (c ShowAgencyRequestRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyRequestRoleType) UnmarshalJSON(b []byte) error {
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
