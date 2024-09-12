package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LayerConfig 图层配置。
type LayerConfig struct {

	// **参数解释**： 图层类型。 **约束限制**： 不涉及。 **取值范围**： * HUMAN:  人物图层 * IMAGE： 素材图片图层 * VIDEO： 素材视频图层 * TEXT: 素材文字图层  **默认取值**： 不涉及
	LayerType LayerConfigLayerType `json:"layer_type"`

	// **参数解释**： 图层所需资产的资产id，外部资产信息无需填写。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	AssetId *string `json:"asset_id,omitempty"`

	// **参数解释**： 多场景素材编组。同一group_id的素材，在应用全局时共享位置信息。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	GroupId *string `json:"group_id,omitempty"`

	Position *LayerPositionConfig `json:"position,omitempty"`

	Size *LayerSizeConfig `json:"size,omitempty"`

	ImageConfig *ImageLayerConfig `json:"image_config,omitempty"`

	VideoConfig *VideoLayerConfig `json:"video_config,omitempty"`

	TextConfig *TextLayerConfig `json:"text_config,omitempty"`
}

func (o LayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayerConfig struct{}"
	}

	return strings.Join([]string{"LayerConfig", string(data)}, " ")
}

type LayerConfigLayerType struct {
	value string
}

type LayerConfigLayerTypeEnum struct {
	HUMAN LayerConfigLayerType
	IMAGE LayerConfigLayerType
	VIDEO LayerConfigLayerType
	TEXT  LayerConfigLayerType
}

func GetLayerConfigLayerTypeEnum() LayerConfigLayerTypeEnum {
	return LayerConfigLayerTypeEnum{
		HUMAN: LayerConfigLayerType{
			value: "HUMAN",
		},
		IMAGE: LayerConfigLayerType{
			value: "IMAGE",
		},
		VIDEO: LayerConfigLayerType{
			value: "VIDEO",
		},
		TEXT: LayerConfigLayerType{
			value: "TEXT",
		},
	}
}

func (c LayerConfigLayerType) Value() string {
	return c.value
}

func (c LayerConfigLayerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LayerConfigLayerType) UnmarshalJSON(b []byte) error {
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
