package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ThresholdInfoResp **参数解释** 监控视图的阈值辅助线配置
type ThresholdInfoResp struct {

	// **参数解释** 监控视图辅助线的阈值 **取值范围** 最小值为0，最大值为2147483647
	Threshold float32 `json:"threshold,omitempty"`

	// **参数解释** 监控视图辅助线的颜色        **取值范围** - \"#B50E65\":紫色 - \"#F23030\":红色 - \"#FF8800\":橙色 - \"#F2E70C\":黄色
	ThresholdColor *ThresholdInfoRespThresholdColor `json:"threshold_color,omitempty"`
}

func (o ThresholdInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThresholdInfoResp struct{}"
	}

	return strings.Join([]string{"ThresholdInfoResp", string(data)}, " ")
}

type ThresholdInfoRespThresholdColor struct {
	value string
}

type ThresholdInfoRespThresholdColorEnum struct {
	B50_E65  ThresholdInfoRespThresholdColor
	F23030   ThresholdInfoRespThresholdColor
	FF8800   ThresholdInfoRespThresholdColor
	F2_E70_C ThresholdInfoRespThresholdColor
}

func GetThresholdInfoRespThresholdColorEnum() ThresholdInfoRespThresholdColorEnum {
	return ThresholdInfoRespThresholdColorEnum{
		B50_E65: ThresholdInfoRespThresholdColor{
			value: "#B50E65",
		},
		F23030: ThresholdInfoRespThresholdColor{
			value: "#F23030",
		},
		FF8800: ThresholdInfoRespThresholdColor{
			value: "#FF8800",
		},
		F2_E70_C: ThresholdInfoRespThresholdColor{
			value: "#F2E70C",
		},
	}
}

func (c ThresholdInfoRespThresholdColor) Value() string {
	return c.value
}

func (c ThresholdInfoRespThresholdColor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThresholdInfoRespThresholdColor) UnmarshalJSON(b []byte) error {
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
