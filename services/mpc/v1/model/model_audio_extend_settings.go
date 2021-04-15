package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AudioExtendSettings struct {
	// 输出策略。  取值如下： - discard - transcode  >- 当视频参数中的“output_policy”为\"discard\"，且音频参数中的“output_policy”为“transcode”时，表示只输出音频。 >- 当视频参数中的“output_policy”为\"transcode\"，且音频参数中的“output_policy”为“discard”时，表示只输出视频。 >- 同时为\"discard\"时不合法。 >- 同时为“transcode”时，表示输出音视频。

	OutputPolicy *AudioExtendSettingsOutputPolicy `json:"output_policy,omitempty"`
	// 声道列表。用来覆盖模板组中的同名字段，如果不配置，则以模板组中的参数为准。

	Channels *[]string `json:"channels,omitempty"`
}

func (o AudioExtendSettings) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AudioExtendSettings struct{}"
	}

	return strings.Join([]string{"AudioExtendSettings", string(data)}, " ")
}

type AudioExtendSettingsOutputPolicy struct {
	value string
}

type AudioExtendSettingsOutputPolicyEnum struct {
	TRANSCODE AudioExtendSettingsOutputPolicy
	DISCARD   AudioExtendSettingsOutputPolicy
	COPY      AudioExtendSettingsOutputPolicy
}

func GetAudioExtendSettingsOutputPolicyEnum() AudioExtendSettingsOutputPolicyEnum {
	return AudioExtendSettingsOutputPolicyEnum{
		TRANSCODE: AudioExtendSettingsOutputPolicy{
			value: "transcode",
		},
		DISCARD: AudioExtendSettingsOutputPolicy{
			value: "discard",
		},
		COPY: AudioExtendSettingsOutputPolicy{
			value: "copy",
		},
	}
}

func (c AudioExtendSettingsOutputPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *AudioExtendSettingsOutputPolicy) UnmarshalJSON(b []byte) error {
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
