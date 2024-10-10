package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LiveRoomInteractionConfig 直播使用互动配置信息
type LiveRoomInteractionConfig struct {

	// 播放类型。 - APPEND：追加，放置在场景播放队列尾部 - INSERT： 插入，在两个音频文件，或者文本句末添加。 - PLAY_NOW : 立即插入，收到指令后，立即播放，无需等待句末。
	PlayType *LiveRoomInteractionConfigPlayType `json:"play_type,omitempty"`

	// 忽略互动回复中断句子后半句不再播放。用于立即中断场景。默认不忽略。
	IgnoreCurrentSentence *bool `json:"ignore_current_sentence,omitempty"`
}

func (o LiveRoomInteractionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveRoomInteractionConfig struct{}"
	}

	return strings.Join([]string{"LiveRoomInteractionConfig", string(data)}, " ")
}

type LiveRoomInteractionConfigPlayType struct {
	value string
}

type LiveRoomInteractionConfigPlayTypeEnum struct {
	APPEND   LiveRoomInteractionConfigPlayType
	INSERT   LiveRoomInteractionConfigPlayType
	PLAY_NOW LiveRoomInteractionConfigPlayType
}

func GetLiveRoomInteractionConfigPlayTypeEnum() LiveRoomInteractionConfigPlayTypeEnum {
	return LiveRoomInteractionConfigPlayTypeEnum{
		APPEND: LiveRoomInteractionConfigPlayType{
			value: "APPEND",
		},
		INSERT: LiveRoomInteractionConfigPlayType{
			value: "INSERT",
		},
		PLAY_NOW: LiveRoomInteractionConfigPlayType{
			value: "PLAY_NOW",
		},
	}
}

func (c LiveRoomInteractionConfigPlayType) Value() string {
	return c.value
}

func (c LiveRoomInteractionConfigPlayType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveRoomInteractionConfigPlayType) UnmarshalJSON(b []byte) error {
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
