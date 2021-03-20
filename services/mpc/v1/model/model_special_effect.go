package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SpecialEffect struct {
	// 特效处理类型。取值： Blur（水印模糊化） Mosaic（水印打马赛克，暂不支持） Replace（替换，暂不支持）

	Type *SpecialEffectType `json:"type,omitempty"`
	// On：表示自动检测特效处理的参数信息，如水印的时间、位置信息等，无需输入参数信息effectinfos OFF：表示需输入特效处理的相关参数信息effectinfos

	AutoDetect *string `json:"auto_detect,omitempty"`
	// 特效处理相关参数，数组 约束：1）每帧视频最多处理2个指定区域；2）每个视频最多处理20个指定区域。

	EffectInfos *[]EffectInfo `json:"effect_infos,omitempty"`
}

func (o SpecialEffect) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SpecialEffect struct{}"
	}

	return strings.Join([]string{"SpecialEffect", string(data)}, " ")
}

type SpecialEffectType struct {
	value string
}

type SpecialEffectTypeEnum struct {
	BLUR    SpecialEffectType
	MOSAIC  SpecialEffectType
	REPLACE SpecialEffectType
}

func GetSpecialEffectTypeEnum() SpecialEffectTypeEnum {
	return SpecialEffectTypeEnum{
		BLUR: SpecialEffectType{
			value: "Blur",
		},
		MOSAIC: SpecialEffectType{
			value: "Mosaic",
		},
		REPLACE: SpecialEffectType{
			value: "Replace",
		},
	}
}

func (c SpecialEffectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SpecialEffectType) UnmarshalJSON(b []byte) error {
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
