package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PlayPolicy 剧本播放策略
type PlayPolicy struct {

	// 剧本重复播放次数。 -1表示持续重复，直至人工停止 0 表示不重复，仅执行一次 其他值n，实际运行次数为n+1次
	RepeatCount *int32 `json:"repeat_count,omitempty"`

	// 是否自动播放剧本。 true: 服务完成任务初始化后，自动播放剧本 false: 服务完成任务初始化后，等待信号后再开始播放剧本
	AutoPlayScript *bool `json:"auto_play_script,omitempty"`

	// 驱动方式。默认TEXT * TEXT: 文本驱动，即通过TTS合成语音 * AUDIO: 语音驱动
	PlayMode *PlayPolicyPlayMode `json:"play_mode,omitempty"`
}

func (o PlayPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlayPolicy struct{}"
	}

	return strings.Join([]string{"PlayPolicy", string(data)}, " ")
}

type PlayPolicyPlayMode struct {
	value string
}

type PlayPolicyPlayModeEnum struct {
	TEXT  PlayPolicyPlayMode
	AUDIO PlayPolicyPlayMode
}

func GetPlayPolicyPlayModeEnum() PlayPolicyPlayModeEnum {
	return PlayPolicyPlayModeEnum{
		TEXT: PlayPolicyPlayMode{
			value: "TEXT",
		},
		AUDIO: PlayPolicyPlayMode{
			value: "AUDIO",
		},
	}
}

func (c PlayPolicyPlayMode) Value() string {
	return c.value
}

func (c PlayPolicyPlayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PlayPolicyPlayMode) UnmarshalJSON(b []byte) error {
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
