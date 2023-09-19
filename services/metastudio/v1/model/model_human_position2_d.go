package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HumanPosition2D 分身数字人在背景图片位置。
type HumanPosition2D struct {

	// 分身数字人在背景图片中的位置。 * LEFT： 左 * MIDDLE： 中 * RIGHT： 右 > 当position_x和position_y参数值存在时，position不生效
	Position *HumanPosition2DPosition `json:"position,omitempty"`

	// 分身数字人X轴位置，即分身数字图片底边中心点像素的X轴的像素值。 > 横屏（16:9）背景图片像素为1920*1080；竖屏（9:16）背景图片像素为1080*1920。
	PositionX *int32 `json:"position_x,omitempty"`

	// 分身数字Y轴位置，即分身数字图片底边中心点像素的Y轴的像素值。 > 横屏（16:9）背景图片像素为1920*1080；竖屏（9:16）背景图片像素为1080*1920。
	PositionY *int32 `json:"position_y,omitempty"`
}

func (o HumanPosition2D) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HumanPosition2D struct{}"
	}

	return strings.Join([]string{"HumanPosition2D", string(data)}, " ")
}

type HumanPosition2DPosition struct {
	value string
}

type HumanPosition2DPositionEnum struct {
	LEFT   HumanPosition2DPosition
	MIDDLE HumanPosition2DPosition
	RIGHT  HumanPosition2DPosition
}

func GetHumanPosition2DPositionEnum() HumanPosition2DPositionEnum {
	return HumanPosition2DPositionEnum{
		LEFT: HumanPosition2DPosition{
			value: "LEFT",
		},
		MIDDLE: HumanPosition2DPosition{
			value: "MIDDLE",
		},
		RIGHT: HumanPosition2DPosition{
			value: "RIGHT",
		},
	}
}

func (c HumanPosition2DPosition) Value() string {
	return c.value
}

func (c HumanPosition2DPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HumanPosition2DPosition) UnmarshalJSON(b []byte) error {
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
