package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesDisplay 显示。
type PoliciesDisplay struct {

	// 显示级别。取值为： LEVEL1：表示等级1。 LEVEL2：表示等级2。 LEVEL3：表示等级3。 LEVEL4：表示等级4（默认/推荐）。 LEVEL5：表示等级5。
	DisplayLevel *PoliciesDisplayDisplayLevel `json:"display_level,omitempty"`

	Options *PoliciesDisplayOptions `json:"options,omitempty"`

	// 是否开启渲染加速。取值为： false：表示关闭。 true：表示开启。
	RenderingAccelerationEnable *bool `json:"rendering_acceleration_enable,omitempty"`

	RenderingAccelerationOptions *PoliciesDisplayRenderingAccelerationOptions `json:"rendering_acceleration_options,omitempty"`

	// 显卡缓存（MB）。取值范围为[0-64]。默认：64。
	VideoCardMemorySize *int32 `json:"video_card_memory_size,omitempty"`

	// 是否开启驱动托管模式。取值为： false：表示关闭。 true：表示开启。
	DriverDelegationModeEnable *bool `json:"driver_delegation_mode_enable,omitempty"`

	// 驱动托管延时（*30ms）。取值范围为[1-100]。默认：80。
	DriverDelegationLatency *int32 `json:"driver_delegation_latency,omitempty"`

	// 驱动托管视频延时（*30ms）。取值范围为[1-100]。默认：80。
	VideoLatency *int32 `json:"video_latency,omitempty"`

	// 计算机修改分辨率：取值为： false：表示关闭。 true：表示开启。
	ChangeResolutionVm *bool `json:"change_resolution_vm,omitempty"`
}

func (o PoliciesDisplay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplay struct{}"
	}

	return strings.Join([]string{"PoliciesDisplay", string(data)}, " ")
}

type PoliciesDisplayDisplayLevel struct {
	value string
}

type PoliciesDisplayDisplayLevelEnum struct {
	LEVEL1 PoliciesDisplayDisplayLevel
	LEVEL2 PoliciesDisplayDisplayLevel
	LEVEL3 PoliciesDisplayDisplayLevel
	LEVEL4 PoliciesDisplayDisplayLevel
	LEVEL5 PoliciesDisplayDisplayLevel
}

func GetPoliciesDisplayDisplayLevelEnum() PoliciesDisplayDisplayLevelEnum {
	return PoliciesDisplayDisplayLevelEnum{
		LEVEL1: PoliciesDisplayDisplayLevel{
			value: "LEVEL1",
		},
		LEVEL2: PoliciesDisplayDisplayLevel{
			value: "LEVEL2",
		},
		LEVEL3: PoliciesDisplayDisplayLevel{
			value: "LEVEL3",
		},
		LEVEL4: PoliciesDisplayDisplayLevel{
			value: "LEVEL4",
		},
		LEVEL5: PoliciesDisplayDisplayLevel{
			value: "LEVEL5",
		},
	}
}

func (c PoliciesDisplayDisplayLevel) Value() string {
	return c.value
}

func (c PoliciesDisplayDisplayLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesDisplayDisplayLevel) UnmarshalJSON(b []byte) error {
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
