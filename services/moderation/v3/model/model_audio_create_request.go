package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 音频内容审核请求体
type AudioCreateRequest struct {
	Data *AudioInputBody `json:"data"`

	// 事件类型，可选值如下： default：默认事件 audiobook：有声书 education：教育音频 game：游戏语音房 live：秀场直播 ecommerce：电商直播 voiceroom：交友语音房 private：私密语音聊天
	EventType AudioCreateRequestEventType `json:"event_type"`

	// 需要检测的风险类型，列表不能为空。 风险类型如下： - porn：涉黄检测 - ad：广告检测 - moan：娇喘检测 - abuse：辱骂检测
	Categories []AudioCreateRequestCategories `json:"categories"`

	// 回调http接口：当该字段非空时，服务将根据该字段回调通知用户审核结果。
	Callback *string `json:"callback,omitempty"`
}

func (o AudioCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioCreateRequest struct{}"
	}

	return strings.Join([]string{"AudioCreateRequest", string(data)}, " ")
}

type AudioCreateRequestEventType struct {
	value string
}

type AudioCreateRequestEventTypeEnum struct {
	DEFAULT   AudioCreateRequestEventType
	AUDIOBOOK AudioCreateRequestEventType
	EDUCATION AudioCreateRequestEventType
	GAME      AudioCreateRequestEventType
	LIVE      AudioCreateRequestEventType
	ECOMMERCE AudioCreateRequestEventType
	VOICEROOM AudioCreateRequestEventType
	PRIVATE   AudioCreateRequestEventType
}

func GetAudioCreateRequestEventTypeEnum() AudioCreateRequestEventTypeEnum {
	return AudioCreateRequestEventTypeEnum{
		DEFAULT: AudioCreateRequestEventType{
			value: "default",
		},
		AUDIOBOOK: AudioCreateRequestEventType{
			value: "audiobook",
		},
		EDUCATION: AudioCreateRequestEventType{
			value: "education",
		},
		GAME: AudioCreateRequestEventType{
			value: "game",
		},
		LIVE: AudioCreateRequestEventType{
			value: "live",
		},
		ECOMMERCE: AudioCreateRequestEventType{
			value: "ecommerce",
		},
		VOICEROOM: AudioCreateRequestEventType{
			value: "voiceroom",
		},
		PRIVATE: AudioCreateRequestEventType{
			value: "private",
		},
	}
}

func (c AudioCreateRequestEventType) Value() string {
	return c.value
}

func (c AudioCreateRequestEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioCreateRequestEventType) UnmarshalJSON(b []byte) error {
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

type AudioCreateRequestCategories struct {
	value string
}

type AudioCreateRequestCategoriesEnum struct {
	POLITICS AudioCreateRequestCategories
	PORN     AudioCreateRequestCategories
	AD       AudioCreateRequestCategories
	MOAN     AudioCreateRequestCategories
	ABUSE    AudioCreateRequestCategories
}

func GetAudioCreateRequestCategoriesEnum() AudioCreateRequestCategoriesEnum {
	return AudioCreateRequestCategoriesEnum{
		POLITICS: AudioCreateRequestCategories{
			value: "politics",
		},
		PORN: AudioCreateRequestCategories{
			value: "porn",
		},
		AD: AudioCreateRequestCategories{
			value: "ad",
		},
		MOAN: AudioCreateRequestCategories{
			value: "moan",
		},
		ABUSE: AudioCreateRequestCategories{
			value: "abuse",
		},
	}
}

func (c AudioCreateRequestCategories) Value() string {
	return c.value
}

func (c AudioCreateRequestCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioCreateRequestCategories) UnmarshalJSON(b []byte) error {
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
