package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 语音合成配置数据结构。
type TtsConfig struct {

	// 语音格式头：wav、mp3、pcm。 默认：wav
	AudioFormat *TtsConfigAudioFormat `json:"audio_format,omitempty"`

	// 采样率：16000、8000 默认：8000
	SampleRate *TtsConfigSampleRate `json:"sample_rate,omitempty"`

	// 语音合成特征字符串，组成形式为{language}_{speaker}_{domain}，即“语种_人员标识_领域”。发音人分为普通发音人和精品发音人。  普通发音人每100字计一次调用，取值范围如下：   chinese_xiaoqi_common  小琪，标准女声发音人。  chinese_xiaoyu_common  小宇，标准男声发音人。  chinese_xiaoyan_common  小燕，温柔女声发音人。  chinese_xiaowang_common  小王，童声发音人。  chinese_xiaowen_common   小雯，柔美女声发音人。  chinese_xiaojing_common 小婧，俏皮女声发音人。  chinese_xiaosong_common  小宋，激昂男声发音人。  chinese_xiaoxia_common   小夏，热情女声发音人。  chinese_xiaodai_common   小呆，呆萌童声发音人。  chinese_xiaoqian_common  小倩，成熟女声发音人。  english_cameal_common    cameal，柔美女声英文发音人。   精品发音人每50字计一次调用，区域仅支持cn-north-4，cn-east-3，暂时不支持音高调节，取值范围如下：  chinese_huaxiaoxia_common  华小夏，热情女声发音人。  chinese_huaxiaogang_common  华晓刚，利落男声发音人。  chinese_huaxiaolu_common  华小璐，知性女声发音人。  chinese_huaxiaoshu_common  华小舒，舒缓女声发音人。  chinese_huaxiaowei_common  华小唯，嗲柔女声发音人。  chinese_huaxiaoliang_common  华小靓，嘹亮女声发音人。  chinese_huaxiaodong_common  华晓东，成熟男声发音人。  chinese_huaxiaoyan_common  华小颜，严厉女声发音人。  chinese_huaxiaoxuan_common  华小萱，台湾女声发音人。  chinese_huaxiaowen_common  华小雯，柔美女声发音人。  chinese_huaxiaoyang_common  华晓阳，朝气男声发音人。  chinese_huaxiaomin_common  华小闽，闽南女声发音人。  chinese_huanvxia_literature 华女侠，武侠女生发音人，只支持16k的采样率。  chinese_huaxiaoxuan_literature 华晓悬，悬疑男声发音人，只支持16k的采样率。  默认：chinese_xiaoyan_common
	Property *TtsConfigProperty `json:"property,omitempty"`

	// 语速。 取值范围：[-500,500]  默认值：0
	Speed *int32 `json:"speed,omitempty"`

	// 音高。 取值范围： [-500,500]  默认值：0
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。 取值范围：[0, 100]  默认值：50
	Volume *int32 `json:"volume,omitempty"`
}

func (o TtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtsConfig struct{}"
	}

	return strings.Join([]string{"TtsConfig", string(data)}, " ")
}

type TtsConfigAudioFormat struct {
	value string
}

type TtsConfigAudioFormatEnum struct {
	WAV TtsConfigAudioFormat
	MP3 TtsConfigAudioFormat
	PCM TtsConfigAudioFormat
}

func GetTtsConfigAudioFormatEnum() TtsConfigAudioFormatEnum {
	return TtsConfigAudioFormatEnum{
		WAV: TtsConfigAudioFormat{
			value: "wav",
		},
		MP3: TtsConfigAudioFormat{
			value: "mp3",
		},
		PCM: TtsConfigAudioFormat{
			value: "pcm",
		},
	}
}

func (c TtsConfigAudioFormat) Value() string {
	return c.value
}

func (c TtsConfigAudioFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TtsConfigAudioFormat) UnmarshalJSON(b []byte) error {
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

type TtsConfigSampleRate struct {
	value string
}

type TtsConfigSampleRateEnum struct {
	E_16000 TtsConfigSampleRate
	E_8000  TtsConfigSampleRate
}

func GetTtsConfigSampleRateEnum() TtsConfigSampleRateEnum {
	return TtsConfigSampleRateEnum{
		E_16000: TtsConfigSampleRate{
			value: "16000",
		},
		E_8000: TtsConfigSampleRate{
			value: "8000",
		},
	}
}

func (c TtsConfigSampleRate) Value() string {
	return c.value
}

func (c TtsConfigSampleRate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TtsConfigSampleRate) UnmarshalJSON(b []byte) error {
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

type TtsConfigProperty struct {
	value string
}

type TtsConfigPropertyEnum struct {
	CHINESE_XIAOQI_COMMON          TtsConfigProperty
	CHINESE_XIAOYU_COMMON          TtsConfigProperty
	CHINESE_XIAOYAN_COMMON         TtsConfigProperty
	CHINESE_XIAOXIA_COMMON         TtsConfigProperty
	CHINESE_XIAODAI_COMMON         TtsConfigProperty
	CHINESE_XIAOQIAN_COMMON        TtsConfigProperty
	CHINESE_XIAOWANG_COMMON        TtsConfigProperty
	CHINESE_XIAOWEN_COMMON         TtsConfigProperty
	CHINESE_XIAOJING_COMMON        TtsConfigProperty
	CHINESE_XIAOSONG_COMMON        TtsConfigProperty
	ENGLISH_CAMEAL_COMMON          TtsConfigProperty
	CHINESE_HUAXIAOXIA_COMMON      TtsConfigProperty
	CHINESE_HUAXIAOGANG_COMMON     TtsConfigProperty
	CHINESE_HUAXIAOLU_COMMON       TtsConfigProperty
	CHINESE_HUAXIAOSHU_COMMON      TtsConfigProperty
	CHINESE_HUAXIAOWEI_COMMON      TtsConfigProperty
	CHINESE_HUAXIAOLIANG_COMMON    TtsConfigProperty
	CHINESE_HUAXIAODONG_COMMON     TtsConfigProperty
	CHINESE_HUAXIAOYAN_COMMON      TtsConfigProperty
	CHINESE_HUAXIAOXUAN_COMMON     TtsConfigProperty
	CHINESE_HUAXIAOWEN_COMMON      TtsConfigProperty
	CHINESE_HUAXIAOYANG_COMMON     TtsConfigProperty
	CHINESE_HUAXIAOMIN_COMMON      TtsConfigProperty
	CHINESE_HUANVXIA_LITERATURE    TtsConfigProperty
	CHINESE_HUAXIAOXUAN_LITERATURE TtsConfigProperty
	CHINESE_HUAXIAOMEI_COMMON      TtsConfigProperty
	CHINESE_HUAXIAOFEI_COMMON      TtsConfigProperty
	CHINESE_HUAXIAOLONG_COMMON     TtsConfigProperty
	CHINESE_HUAXIAORUI_COMMON      TtsConfigProperty
}

func GetTtsConfigPropertyEnum() TtsConfigPropertyEnum {
	return TtsConfigPropertyEnum{
		CHINESE_XIAOQI_COMMON: TtsConfigProperty{
			value: "chinese_xiaoqi_common",
		},
		CHINESE_XIAOYU_COMMON: TtsConfigProperty{
			value: "chinese_xiaoyu_common",
		},
		CHINESE_XIAOYAN_COMMON: TtsConfigProperty{
			value: "chinese_xiaoyan_common",
		},
		CHINESE_XIAOXIA_COMMON: TtsConfigProperty{
			value: "chinese_xiaoxia_common",
		},
		CHINESE_XIAODAI_COMMON: TtsConfigProperty{
			value: "chinese_xiaodai_common",
		},
		CHINESE_XIAOQIAN_COMMON: TtsConfigProperty{
			value: "chinese_xiaoqian_common",
		},
		CHINESE_XIAOWANG_COMMON: TtsConfigProperty{
			value: "chinese_xiaowang_common",
		},
		CHINESE_XIAOWEN_COMMON: TtsConfigProperty{
			value: "chinese_xiaowen_common",
		},
		CHINESE_XIAOJING_COMMON: TtsConfigProperty{
			value: "chinese_xiaojing_common",
		},
		CHINESE_XIAOSONG_COMMON: TtsConfigProperty{
			value: "chinese_xiaosong_common",
		},
		ENGLISH_CAMEAL_COMMON: TtsConfigProperty{
			value: "english_cameal_common",
		},
		CHINESE_HUAXIAOXIA_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaoxia_common",
		},
		CHINESE_HUAXIAOGANG_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaogang_common",
		},
		CHINESE_HUAXIAOLU_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaolu_common",
		},
		CHINESE_HUAXIAOSHU_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaoshu_common",
		},
		CHINESE_HUAXIAOWEI_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaowei_common",
		},
		CHINESE_HUAXIAOLIANG_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaoliang_common",
		},
		CHINESE_HUAXIAODONG_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaodong_common",
		},
		CHINESE_HUAXIAOYAN_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaoyan_common",
		},
		CHINESE_HUAXIAOXUAN_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaoxuan_common",
		},
		CHINESE_HUAXIAOWEN_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaowen_common",
		},
		CHINESE_HUAXIAOYANG_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaoyang_common",
		},
		CHINESE_HUAXIAOMIN_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaomin_common",
		},
		CHINESE_HUANVXIA_LITERATURE: TtsConfigProperty{
			value: "chinese_huanvxia_literature",
		},
		CHINESE_HUAXIAOXUAN_LITERATURE: TtsConfigProperty{
			value: "chinese_huaxiaoxuan_literature",
		},
		CHINESE_HUAXIAOMEI_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaomei_common",
		},
		CHINESE_HUAXIAOFEI_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaofei_common",
		},
		CHINESE_HUAXIAOLONG_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaolong_common",
		},
		CHINESE_HUAXIAORUI_COMMON: TtsConfigProperty{
			value: "chinese_huaxiaorui_common",
		},
	}
}

func (c TtsConfigProperty) Value() string {
	return c.value
}

func (c TtsConfigProperty) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TtsConfigProperty) UnmarshalJSON(b []byte) error {
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
