package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateSubnetBandwidthControlListReq struct {

	// 控制模式 - BLACK：黑名单控制。 - WHITE：白名单控制。
	ControlMode *UpdateSubnetBandwidthControlListReqControlMode `json:"control_mode,omitempty"`

	// 待添加的桌面。
	AddControlList *[]string `json:"add_control_list,omitempty"`

	// 待删除的桌面。
	RemoveControlList *[]string `json:"remove_control_list,omitempty"`
}

func (o UpdateSubnetBandwidthControlListReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetBandwidthControlListReq struct{}"
	}

	return strings.Join([]string{"UpdateSubnetBandwidthControlListReq", string(data)}, " ")
}

type UpdateSubnetBandwidthControlListReqControlMode struct {
	value string
}

type UpdateSubnetBandwidthControlListReqControlModeEnum struct {
	BLACK UpdateSubnetBandwidthControlListReqControlMode
	WHITE UpdateSubnetBandwidthControlListReqControlMode
}

func GetUpdateSubnetBandwidthControlListReqControlModeEnum() UpdateSubnetBandwidthControlListReqControlModeEnum {
	return UpdateSubnetBandwidthControlListReqControlModeEnum{
		BLACK: UpdateSubnetBandwidthControlListReqControlMode{
			value: "BLACK",
		},
		WHITE: UpdateSubnetBandwidthControlListReqControlMode{
			value: "WHITE",
		},
	}
}

func (c UpdateSubnetBandwidthControlListReqControlMode) Value() string {
	return c.value
}

func (c UpdateSubnetBandwidthControlListReqControlMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubnetBandwidthControlListReqControlMode) UnmarshalJSON(b []byte) error {
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
