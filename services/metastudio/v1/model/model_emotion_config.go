package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EmotionConfig 情感标签配置。
type EmotionConfig struct {

	// 情感标签配置。 * HAPPY：开心 * SAD：悲伤 * CALM：平静 * ANGER：愤怒  默认HAPPY。
	Emotion *EmotionConfigEmotion `json:"emotion,omitempty"`
}

func (o EmotionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmotionConfig struct{}"
	}

	return strings.Join([]string{"EmotionConfig", string(data)}, " ")
}

type EmotionConfigEmotion struct {
	value string
}

type EmotionConfigEmotionEnum struct {
	HAPPY EmotionConfigEmotion
	SAD   EmotionConfigEmotion
	CALM  EmotionConfigEmotion
	ANGER EmotionConfigEmotion
}

func GetEmotionConfigEmotionEnum() EmotionConfigEmotionEnum {
	return EmotionConfigEmotionEnum{
		HAPPY: EmotionConfigEmotion{
			value: "HAPPY",
		},
		SAD: EmotionConfigEmotion{
			value: "SAD",
		},
		CALM: EmotionConfigEmotion{
			value: "CALM",
		},
		ANGER: EmotionConfigEmotion{
			value: "ANGER",
		},
	}
}

func (c EmotionConfigEmotion) Value() string {
	return c.value
}

func (c EmotionConfigEmotion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EmotionConfigEmotion) UnmarshalJSON(b []byte) error {
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
