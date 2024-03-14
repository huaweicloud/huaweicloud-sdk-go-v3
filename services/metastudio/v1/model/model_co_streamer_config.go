package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CoStreamerConfig 助播配置
type CoStreamerConfig struct {
	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	// 助播出声时主播行为。 * SILENCE：静默 * VOLUME_DOWN：音量降低
	StreamerAction *CoStreamerConfigStreamerAction `json:"streamer_action,omitempty"`
}

func (o CoStreamerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoStreamerConfig struct{}"
	}

	return strings.Join([]string{"CoStreamerConfig", string(data)}, " ")
}

type CoStreamerConfigStreamerAction struct {
	value string
}

type CoStreamerConfigStreamerActionEnum struct {
	SILENCE     CoStreamerConfigStreamerAction
	VOLUME_DOWN CoStreamerConfigStreamerAction
}

func GetCoStreamerConfigStreamerActionEnum() CoStreamerConfigStreamerActionEnum {
	return CoStreamerConfigStreamerActionEnum{
		SILENCE: CoStreamerConfigStreamerAction{
			value: "SILENCE",
		},
		VOLUME_DOWN: CoStreamerConfigStreamerAction{
			value: "VOLUME_DOWN",
		},
	}
}

func (c CoStreamerConfigStreamerAction) Value() string {
	return c.value
}

func (c CoStreamerConfigStreamerAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CoStreamerConfigStreamerAction) UnmarshalJSON(b []byte) error {
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
