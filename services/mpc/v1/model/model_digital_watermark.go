package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type DigitalWatermark struct {
	// 数字水印属性。 - ROBUST：水印鲁棒性最高 - MEDIUM：水印鲁棒性和视频质量折中(默认，暂时只支持该选项) - QUALITY：视频质量最好 - MEZZ：具有最高感官质量的水印 - CAMCORDING：最强大的水印配置文件，支持摄录攻击
	Profile DigitalWatermarkProfile `json:"profile"`
}

func (o DigitalWatermark) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DigitalWatermark struct{}"
	}

	return strings.Join([]string{"DigitalWatermark", string(data)}, " ")
}

type DigitalWatermarkProfile struct {
	value string
}

type DigitalWatermarkProfileEnum struct {
	ROBUST     DigitalWatermarkProfile
	MEDIUM     DigitalWatermarkProfile
	QUALITY    DigitalWatermarkProfile
	MEZZ       DigitalWatermarkProfile
	CAMCORDING DigitalWatermarkProfile
}

func GetDigitalWatermarkProfileEnum() DigitalWatermarkProfileEnum {
	return DigitalWatermarkProfileEnum{
		ROBUST: DigitalWatermarkProfile{
			value: "ROBUST",
		},
		MEDIUM: DigitalWatermarkProfile{
			value: "MEDIUM",
		},
		QUALITY: DigitalWatermarkProfile{
			value: "QUALITY",
		},
		MEZZ: DigitalWatermarkProfile{
			value: "MEZZ",
		},
		CAMCORDING: DigitalWatermarkProfile{
			value: "CAMCORDING",
		},
	}
}

func (c DigitalWatermarkProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DigitalWatermarkProfile) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
