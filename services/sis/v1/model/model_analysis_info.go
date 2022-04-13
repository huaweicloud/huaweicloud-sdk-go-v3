package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type AnalysisInfo struct {
	// 是否需要做话者分离。缺省为true，表示会进行话者分离，识别结果中会包含role项（角色）。如果diarization为false, 那么结果中不会出现role项。

	Diarization *bool `json:"diarization,omitempty"`
	// 语音文件声道信息，可以为MONO（缺省), LEFT_AGENT, RIGHT_AGENT。  如果channel 为MONO，那么原始文件需要为单声道文件。  如果为双声道文件，系统会将其转换成单声道文件，可能会影响识别效果。  如果 channel 为 LEFT_AGENT或RIGHT_AGENT, 则原始文件需要为双声道文件，如果为单声道文件，系统会将其转换成双声道文件，可能会影响识别效果。  当channel 为 LEFT_AGENT或RIGHT_AGENT，且diarization为true时，系统会按照配置给出对应角色。其中：  LEFT_AGENT 指定左声道语音为agent（坐席）,  RIGHT_AGENT 指定右声道为agent（坐席）。

	Channel *AnalysisInfoChannel `json:"channel,omitempty"`
	// 是否需要做情绪检测, 缺省为true。

	Emotion *bool `json:"emotion,omitempty"`
	// 是否需要输出语速信息, 缺省为true。

	Speed *bool `json:"speed,omitempty"`
}

func (o AnalysisInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisInfo struct{}"
	}

	return strings.Join([]string{"AnalysisInfo", string(data)}, " ")
}

type AnalysisInfoChannel struct {
	value string
}

type AnalysisInfoChannelEnum struct {
	MONO        AnalysisInfoChannel
	LEFT_AGENT  AnalysisInfoChannel
	RIGHT_AGENT AnalysisInfoChannel
}

func GetAnalysisInfoChannelEnum() AnalysisInfoChannelEnum {
	return AnalysisInfoChannelEnum{
		MONO: AnalysisInfoChannel{
			value: "MONO",
		},
		LEFT_AGENT: AnalysisInfoChannel{
			value: "LEFT_AGENT",
		},
		RIGHT_AGENT: AnalysisInfoChannel{
			value: "RIGHT_AGENT",
		},
	}
}

func (c AnalysisInfoChannel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnalysisInfoChannel) UnmarshalJSON(b []byte) error {
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
