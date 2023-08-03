package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AnimationAssetMeta 动作动画资产元数据。
type AnimationAssetMeta struct {

	// 数字人模型风格ID。
	StyleId *string `json:"style_id,omitempty"`

	// 动作动画时长。
	Duration *float32 `json:"duration,omitempty"`

	// 动作是否需要自动解析。
	AutoAnalysis *bool `json:"auto_analysis,omitempty"`

	// 语音延迟播放时长。  单位秒。  使用场景举例：入场动画3秒，voice_delay设置成4秒，则语音从入场动画开始后第4秒开始播放。
	VoiceDelay *float32 `json:"voice_delay,omitempty"`

	// 动画插入位置限制。 * ONLY_BEGINNING：视频制作时，动画只允许出现在起始位置。 * ONLY_END：视频制作时，动画只允许出现在结束位置。
	AnimationInsertRestriction *AnimationAssetMetaAnimationInsertRestriction `json:"animation_insert_restriction,omitempty"`
}

func (o AnimationAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnimationAssetMeta struct{}"
	}

	return strings.Join([]string{"AnimationAssetMeta", string(data)}, " ")
}

type AnimationAssetMetaAnimationInsertRestriction struct {
	value string
}

type AnimationAssetMetaAnimationInsertRestrictionEnum struct {
	ONLY_BEGINNING AnimationAssetMetaAnimationInsertRestriction
	ONLY_END       AnimationAssetMetaAnimationInsertRestriction
}

func GetAnimationAssetMetaAnimationInsertRestrictionEnum() AnimationAssetMetaAnimationInsertRestrictionEnum {
	return AnimationAssetMetaAnimationInsertRestrictionEnum{
		ONLY_BEGINNING: AnimationAssetMetaAnimationInsertRestriction{
			value: "ONLY_BEGINNING",
		},
		ONLY_END: AnimationAssetMetaAnimationInsertRestriction{
			value: "ONLY_END",
		},
	}
}

func (c AnimationAssetMetaAnimationInsertRestriction) Value() string {
	return c.value
}

func (c AnimationAssetMetaAnimationInsertRestriction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnimationAssetMetaAnimationInsertRestriction) UnmarshalJSON(b []byte) error {
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
