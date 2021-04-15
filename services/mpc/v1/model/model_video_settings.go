package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoSettings struct {
	// 帧率转换开关。取值范围： - ENABLED - DISABLED

	Frc *VideoSettingsFrc `json:"frc,omitempty"`
}

func (o VideoSettings) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VideoSettings struct{}"
	}

	return strings.Join([]string{"VideoSettings", string(data)}, " ")
}

type VideoSettingsFrc struct {
	value string
}

type VideoSettingsFrcEnum struct {
	ENABLED  VideoSettingsFrc
	DISABLED VideoSettingsFrc
}

func GetVideoSettingsFrcEnum() VideoSettingsFrcEnum {
	return VideoSettingsFrcEnum{
		ENABLED: VideoSettingsFrc{
			value: "ENABLED",
		},
		DISABLED: VideoSettingsFrc{
			value: "DISABLED",
		},
	}
}

func (c VideoSettingsFrc) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VideoSettingsFrc) UnmarshalJSON(b []byte) error {
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
