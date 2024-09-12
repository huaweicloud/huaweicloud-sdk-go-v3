package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmartTextLayerConfig 素材视频图层配置。
type SmartTextLayerConfig struct {

	// **参数解释**： 文本类型。 * DYNAMIC：动态文本，需要进行关键字替换。 * STATIC：静态文本。
	TextType *SmartTextLayerConfigTextType `json:"text_type,omitempty"`

	// 文本。
	TextContext *string `json:"text_context,omitempty"`

	// **参数解释**： 字体。当前支持的字体： * HarmonyOS_Sans_SC_Black：鸿蒙粗体 * HarmonyOS_Sans_SC_Regular：鸿蒙常规 * HarmonyOS_Sans_SC_Thin：鸿蒙细体 * fzyouh：方正瘦体
	FontName *string `json:"font_name,omitempty"`

	// **参数解释**： 字体大小（像素）。  业务取值范围：[4, 120]，请以业务取值范围为准。
	FontSize *int32 `json:"font_size,omitempty"`

	// **参数解释**： 字体颜色。RGB颜色值。
	FontColor *string `json:"font_color,omitempty"`

	// **参数解释**： 文本显示时长，单位s。 显示时长规则为，若携带reply_texts、reply_audios，则与播放语音内容时长保持一致；若未携带，则与匹配的关键词语音内容时长保持一致。
	DisplayDuration *int32 `json:"display_duration,omitempty"`
}

func (o SmartTextLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartTextLayerConfig struct{}"
	}

	return strings.Join([]string{"SmartTextLayerConfig", string(data)}, " ")
}

type SmartTextLayerConfigTextType struct {
	value string
}

type SmartTextLayerConfigTextTypeEnum struct {
	DYNAMIC SmartTextLayerConfigTextType
	STATIC  SmartTextLayerConfigTextType
}

func GetSmartTextLayerConfigTextTypeEnum() SmartTextLayerConfigTextTypeEnum {
	return SmartTextLayerConfigTextTypeEnum{
		DYNAMIC: SmartTextLayerConfigTextType{
			value: "DYNAMIC",
		},
		STATIC: SmartTextLayerConfigTextType{
			value: "STATIC",
		},
	}
}

func (c SmartTextLayerConfigTextType) Value() string {
	return c.value
}

func (c SmartTextLayerConfigTextType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartTextLayerConfigTextType) UnmarshalJSON(b []byte) error {
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
