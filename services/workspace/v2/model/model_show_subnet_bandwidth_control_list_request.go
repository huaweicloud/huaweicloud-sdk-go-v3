package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSubnetBandwidthControlListRequest Request Object
type ShowSubnetBandwidthControlListRequest struct {

	// 云办公带宽id。
	BandwidthId string `json:"bandwidth_id"`

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面分配用户。
	UserName *string `json:"user_name,omitempty"`

	// 控制模式，支持： - BLACK 黑名单 - WHITE 白名单
	ControlMode *ShowSubnetBandwidthControlListRequestControlMode `json:"control_mode,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。如果不指定，则返回所有符合条件的记录。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowSubnetBandwidthControlListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubnetBandwidthControlListRequest struct{}"
	}

	return strings.Join([]string{"ShowSubnetBandwidthControlListRequest", string(data)}, " ")
}

type ShowSubnetBandwidthControlListRequestControlMode struct {
	value string
}

type ShowSubnetBandwidthControlListRequestControlModeEnum struct {
	BLACK ShowSubnetBandwidthControlListRequestControlMode
	WHITE ShowSubnetBandwidthControlListRequestControlMode
}

func GetShowSubnetBandwidthControlListRequestControlModeEnum() ShowSubnetBandwidthControlListRequestControlModeEnum {
	return ShowSubnetBandwidthControlListRequestControlModeEnum{
		BLACK: ShowSubnetBandwidthControlListRequestControlMode{
			value: "BLACK",
		},
		WHITE: ShowSubnetBandwidthControlListRequestControlMode{
			value: "WHITE",
		},
	}
}

func (c ShowSubnetBandwidthControlListRequestControlMode) Value() string {
	return c.value
}

func (c ShowSubnetBandwidthControlListRequestControlMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSubnetBandwidthControlListRequestControlMode) UnmarshalJSON(b []byte) error {
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
