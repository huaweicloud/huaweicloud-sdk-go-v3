package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LiveWarningItem 开播风险警告
type LiveWarningItem struct {

	// 告警类型。 - TOO_LESSS_SCRIPT_ITEMS：段落（话术）数量太少。 - TOO_SHORT_SCRIPT_TIME：段落（话术）总时长太短。 - TOO_LESS_DANMAKU_RULES： 弹幕互动规则太少。 - RANDOM_PLAY_CLOSED: 随机播放开关关闭。 - ROTATION_MODEL_CLOSED: 主播轮转未配置。
	WarningType *LiveWarningItemWarningType `json:"warning_type,omitempty"`
}

func (o LiveWarningItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveWarningItem struct{}"
	}

	return strings.Join([]string{"LiveWarningItem", string(data)}, " ")
}

type LiveWarningItemWarningType struct {
	value string
}

type LiveWarningItemWarningTypeEnum struct {
	TOO_LESSS_SCRIPT_ITEMS LiveWarningItemWarningType
	TOO_SHORT_SCRIPT_TIME  LiveWarningItemWarningType
	TOO_LESS_DANMAKU_RULES LiveWarningItemWarningType
	RANDOM_PLAY_CLOSED     LiveWarningItemWarningType
	ROTATION_MODEL_CLOSED  LiveWarningItemWarningType
}

func GetLiveWarningItemWarningTypeEnum() LiveWarningItemWarningTypeEnum {
	return LiveWarningItemWarningTypeEnum{
		TOO_LESSS_SCRIPT_ITEMS: LiveWarningItemWarningType{
			value: "TOO_LESSS_SCRIPT_ITEMS",
		},
		TOO_SHORT_SCRIPT_TIME: LiveWarningItemWarningType{
			value: "TOO_SHORT_SCRIPT_TIME",
		},
		TOO_LESS_DANMAKU_RULES: LiveWarningItemWarningType{
			value: "TOO_LESS_DANMAKU_RULES",
		},
		RANDOM_PLAY_CLOSED: LiveWarningItemWarningType{
			value: "RANDOM_PLAY_CLOSED",
		},
		ROTATION_MODEL_CLOSED: LiveWarningItemWarningType{
			value: "ROTATION_MODEL_CLOSED",
		},
	}
}

func (c LiveWarningItemWarningType) Value() string {
	return c.value
}

func (c LiveWarningItemWarningType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveWarningItemWarningType) UnmarshalJSON(b []byte) error {
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
