package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoExtendSettings struct {
	// 扩展编码质量等级，用于覆盖模板中的同名参数。取值如下： - SPEED - HIGHQUALITY

	Preset *VideoExtendSettingsPreset `json:"preset,omitempty"`
}

func (o VideoExtendSettings) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VideoExtendSettings struct{}"
	}

	return strings.Join([]string{"VideoExtendSettings", string(data)}, " ")
}

type VideoExtendSettingsPreset struct {
	value string
}

type VideoExtendSettingsPresetEnum struct {
	SPEED       VideoExtendSettingsPreset
	HIGHQUALITY VideoExtendSettingsPreset
}

func GetVideoExtendSettingsPresetEnum() VideoExtendSettingsPresetEnum {
	return VideoExtendSettingsPresetEnum{
		SPEED: VideoExtendSettingsPreset{
			value: "SPEED",
		},
		HIGHQUALITY: VideoExtendSettingsPreset{
			value: "HIGHQUALITY",
		},
	}
}

func (c VideoExtendSettingsPreset) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VideoExtendSettingsPreset) UnmarshalJSON(b []byte) error {
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
