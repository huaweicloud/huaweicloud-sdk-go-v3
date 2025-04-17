package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ThresholdInfo 监控视图的阈值辅助线配置
type ThresholdInfo struct {

	// 监控视图辅助线的阈值
	Threshold float32 `json:"threshold"`

	// 监控视图辅助线的颜色,\"#B50E65\"紫色,\"#F23030\"红色,\"#FF8800\"橙色,\"#F2E70C\"黄色
	ThresholdColor ThresholdInfoThresholdColor `json:"threshold_color"`
}

func (o ThresholdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThresholdInfo struct{}"
	}

	return strings.Join([]string{"ThresholdInfo", string(data)}, " ")
}

type ThresholdInfoThresholdColor struct {
	value string
}

type ThresholdInfoThresholdColorEnum struct {
	B50_E65  ThresholdInfoThresholdColor
	F23030   ThresholdInfoThresholdColor
	FF8800   ThresholdInfoThresholdColor
	F2_E70_C ThresholdInfoThresholdColor
}

func GetThresholdInfoThresholdColorEnum() ThresholdInfoThresholdColorEnum {
	return ThresholdInfoThresholdColorEnum{
		B50_E65: ThresholdInfoThresholdColor{
			value: "#B50E65",
		},
		F23030: ThresholdInfoThresholdColor{
			value: "#F23030",
		},
		FF8800: ThresholdInfoThresholdColor{
			value: "#FF8800",
		},
		F2_E70_C: ThresholdInfoThresholdColor{
			value: "#F2E70C",
		},
	}
}

func (c ThresholdInfoThresholdColor) Value() string {
	return c.value
}

func (c ThresholdInfoThresholdColor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThresholdInfoThresholdColor) UnmarshalJSON(b []byte) error {
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
