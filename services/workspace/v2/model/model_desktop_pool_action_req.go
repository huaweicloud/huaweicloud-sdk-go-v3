package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DesktopPoolActionReq 操作桌面池请求。
type DesktopPoolActionReq struct {

	// 操作类型。 -os-start 开机。 -os-stop 关机。 -reboot 重启。 -hibernate 休眠。
	OpType string `json:"op_type"`

	// 执行类型。例如type为HARD，op_type为os-stop代表强制关机。 - SOFT：普通操作。 - HARD：强制操作。
	Type *DesktopPoolActionReqType `json:"type,omitempty"`
}

func (o DesktopPoolActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopPoolActionReq struct{}"
	}

	return strings.Join([]string{"DesktopPoolActionReq", string(data)}, " ")
}

type DesktopPoolActionReqType struct {
	value string
}

type DesktopPoolActionReqTypeEnum struct {
	SOFT DesktopPoolActionReqType
	HARD DesktopPoolActionReqType
}

func GetDesktopPoolActionReqTypeEnum() DesktopPoolActionReqTypeEnum {
	return DesktopPoolActionReqTypeEnum{
		SOFT: DesktopPoolActionReqType{
			value: "SOFT",
		},
		HARD: DesktopPoolActionReqType{
			value: "HARD",
		},
	}
}

func (c DesktopPoolActionReqType) Value() string {
	return c.value
}

func (c DesktopPoolActionReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DesktopPoolActionReqType) UnmarshalJSON(b []byte) error {
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
