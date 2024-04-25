package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AudioCreateRequest 音频内容审核请求体
type AudioCreateRequest struct {
	Data *AudioInputBody `json:"data"`

	// 用户在控制台界面创建的biz_type名称，如果请求参数中传了biz_type则优先使用biz_type；如果用户没传biz_type则event_type和categories必须传。
	BizType *string `json:"biz_type,omitempty"`

	// 事件类型，可选值如下： default：默认事件 audiobook：有声书 education：教育音频 game：游戏语音房 live：秀场直播 ecommerce：电商直播 voiceroom：交友语音房 private：私密语音聊天
	EventType *AudioCreateRequestEventType `json:"event_type,omitempty"`

	// 需要检测的风险类型，列表不能为空。 风险类型如下： - porn：涉黄检测 - ad：广告检测 - moan：娇喘检测 - abuse：辱骂检测
	Categories *[]AudioCreateRequestCategories `json:"categories,omitempty"`

	// 回调http接口：当该字段非空时，服务将根据该字段回调通知用户审核结果。
	Callback *string `json:"callback,omitempty"`

	// 用于回调通知时校验请求由华为云内容安全服务发起，由您自定义。随机字符串，由英文字母、数字、下划线组成，不超过64个字符。 说明：当seed非空时，headers中将包含X-Auth-Signature字段，字段的值使用HmacSHA256算法生成，待加密字符串由create_time、job_id、request_id、seed按照顺序拼接而成，密钥为seed。
	Seed *string `json:"seed,omitempty"`
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
