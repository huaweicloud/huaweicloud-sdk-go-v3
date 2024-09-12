package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PlayPolicy 剧本播放策略
type PlayPolicy struct {

	// **参数解释**： 剧本重复播放次数。 * -1：表示持续重复，直至人工停止。 * 0：表示不重复，仅执行一次。 * 其他值n：实际运行次数为n+1次。  **约束限制**： 不涉及。
	RepeatCount *int32 `json:"repeat_count,omitempty"`

	// **参数解释**： 是否自动播放剧本。 **约束限制**： 不涉及。 **取值范围**： * true：服务完成任务初始化后，自动播放剧本。 * false：服务完成任务初始化后，等待信号后再开始播放剧本。
	AutoPlayScript *bool `json:"auto_play_script,omitempty"`

	// **参数解释**： 驱动方式。 **约束限制**： 不涉及。 **取值范围**： * TEXT：文本驱动，即通过TTS合成语音。 * AUDIO：语音驱动。 * NO_PRESET：无预置剧本，人工控制模式。
	PlayMode *PlayPolicyPlayMode `json:"play_mode,omitempty"`

	// **参数解释**： 随机播报模式。 **约束限制**： 从第二轮播报开始随机。 **取值范围**： * NONE：不启动随机播报。 * SCENE：按场景随机播报。场景内段落按顺序播报。 * SCRIPT_ITEM：按段落随机播报。场景按顺序播报。 * SCENE_AND_SCRIPT_ITEM：场景和段落都随机播报。
	RandomPlayMode *PlayPolicyRandomPlayMode `json:"random_play_mode,omitempty"`

	// **参数解释**： 是否需要独立采集端。用于客户端播放与命令分离场景。 **约束限制**： 不涉及。 **取值范围**： * true：分配CAPTURE、PLAYER两个RTC用户。 * fasle：仅分配PLAYER一个RTC用户。
	NeedIndependentCaptureClient *bool `json:"need_independent_capture_client,omitempty"`
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
	TEXT      PlayPolicyPlayMode
	AUDIO     PlayPolicyPlayMode
	NO_PRESET PlayPolicyPlayMode
}

func GetPlayPolicyPlayModeEnum() PlayPolicyPlayModeEnum {
	return PlayPolicyPlayModeEnum{
		TEXT: PlayPolicyPlayMode{
			value: "TEXT",
		},
		AUDIO: PlayPolicyPlayMode{
			value: "AUDIO",
		},
		NO_PRESET: PlayPolicyPlayMode{
			value: "NO_PRESET",
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

type PlayPolicyRandomPlayMode struct {
	value string
}

type PlayPolicyRandomPlayModeEnum struct {
	NONE                  PlayPolicyRandomPlayMode
	SCENE                 PlayPolicyRandomPlayMode
	SCRIPT_ITEM           PlayPolicyRandomPlayMode
	SCENE_AND_SCRIPT_ITEM PlayPolicyRandomPlayMode
}

func GetPlayPolicyRandomPlayModeEnum() PlayPolicyRandomPlayModeEnum {
	return PlayPolicyRandomPlayModeEnum{
		NONE: PlayPolicyRandomPlayMode{
			value: "NONE",
		},
		SCENE: PlayPolicyRandomPlayMode{
			value: "SCENE",
		},
		SCRIPT_ITEM: PlayPolicyRandomPlayMode{
			value: "SCRIPT_ITEM",
		},
		SCENE_AND_SCRIPT_ITEM: PlayPolicyRandomPlayMode{
			value: "SCENE_AND_SCRIPT_ITEM",
		},
	}
}

func (c PlayPolicyRandomPlayMode) Value() string {
	return c.value
}

func (c PlayPolicyRandomPlayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PlayPolicyRandomPlayMode) UnmarshalJSON(b []byte) error {
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
