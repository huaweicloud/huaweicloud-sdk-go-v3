package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WatermarkLocation struct {

	// 水印位置。  包含如下选项： - TOPLEFT：左上 - TOPRIGHT：右上 - BOTTOMLEFT：左下 - BOTTOMRIGHT：右下 - RANDOM：随机模式，图片水印将随机在视频流的左上、右上、左下、右下四个位置展示。
	Location *WatermarkLocationLocation `json:"location,omitempty"`

	// 水印相对输出视频的水平偏移量。 说明：值有两种形式： 1）整数型代表偏移像素，范围[1，4096]，单位px。 2）小数型代表水平偏移量与输出分辨率宽的比率，范围[0，1) 百分比限制最多保留小数点后4位
	XOffset float32 `json:"x_offset,omitempty"`

	// 水印相对输出视频的垂直偏移量 说明: 值有两种形式： 1）整数型代表偏移像素，范围[1，4096]，单位px。 2）小数型代表垂直偏移量与输出分辨率高的比率，范围[0，1) 百分比限制最多保留小数点后4位
	YOffset float32 `json:"y_offset,omitempty"`
}

func (o WatermarkLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WatermarkLocation struct{}"
	}

	return strings.Join([]string{"WatermarkLocation", string(data)}, " ")
}

type WatermarkLocationLocation struct {
	value string
}

type WatermarkLocationLocationEnum struct {
	TOPLEFT     WatermarkLocationLocation
	TOPRIGHT    WatermarkLocationLocation
	BOTTOMLEFT  WatermarkLocationLocation
	BOTTOMRIGHT WatermarkLocationLocation
	RANDOM      WatermarkLocationLocation
}

func GetWatermarkLocationLocationEnum() WatermarkLocationLocationEnum {
	return WatermarkLocationLocationEnum{
		TOPLEFT: WatermarkLocationLocation{
			value: "TOPLEFT",
		},
		TOPRIGHT: WatermarkLocationLocation{
			value: "TOPRIGHT",
		},
		BOTTOMLEFT: WatermarkLocationLocation{
			value: "BOTTOMLEFT",
		},
		BOTTOMRIGHT: WatermarkLocationLocation{
			value: "BOTTOMRIGHT",
		},
		RANDOM: WatermarkLocationLocation{
			value: "RANDOM",
		},
	}
}

func (c WatermarkLocationLocation) Value() string {
	return c.value
}

func (c WatermarkLocationLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WatermarkLocationLocation) UnmarshalJSON(b []byte) error {
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
