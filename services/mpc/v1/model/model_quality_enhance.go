package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type QualityEnhance struct {
	// 针对一般质量、无明显问题的普通片源，通过增强、锐化等技术明显提升主观效果。单纯该处理操作前后，分辨率、帧率等参数不发生变化。 可和old_repair、super_resolution、super_framerate、SDRToHDR组合使用。
	NormalEnhance *QualityEnhanceNormalEnhance `json:"normal_enhance,omitempty"`
	// 针对旧片、老片，画质主观质量比较低的片源，通过降噪、去压缩失真等视频增强技术，提升画质主观效果。
	Revive *QualityEnhanceRevive `json:"revive,omitempty"`
	// 超动态范围，提升视频动态范围，明显提升片源动态范围。单纯该处理操作前后，分辨率、帧率等参数不发生变化，动态范围、色域范围、码率发生变化。 可和normal_ enhance组合使用。 取值范围： - SDRtoHDR10 ：转换模式1，为标准模式 - SDRtoHDRFLAT”：转换模式2，清新模式，基本不改变源片的饱和度，适用于饱和度高的SDR源片转换为HDR
	SdrToHdr *QualityEnhanceSdrToHdr `json:"sdr_to_hdr,omitempty"`
}

func (o QualityEnhance) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QualityEnhance struct{}"
	}

	return strings.Join([]string{"QualityEnhance", string(data)}, " ")
}

type QualityEnhanceNormalEnhance struct {
	value string
}

type QualityEnhanceNormalEnhanceEnum struct {
	NORMAL QualityEnhanceNormalEnhance
}

func GetQualityEnhanceNormalEnhanceEnum() QualityEnhanceNormalEnhanceEnum {
	return QualityEnhanceNormalEnhanceEnum{
		NORMAL: QualityEnhanceNormalEnhance{
			value: "normal",
		},
	}
}

func (c QualityEnhanceNormalEnhance) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QualityEnhanceNormalEnhance) UnmarshalJSON(b []byte) error {
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

type QualityEnhanceRevive struct {
	value string
}

type QualityEnhanceReviveEnum struct {
	NORMAL QualityEnhanceRevive
}

func GetQualityEnhanceReviveEnum() QualityEnhanceReviveEnum {
	return QualityEnhanceReviveEnum{
		NORMAL: QualityEnhanceRevive{
			value: "normal",
		},
	}
}

func (c QualityEnhanceRevive) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QualityEnhanceRevive) UnmarshalJSON(b []byte) error {
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

type QualityEnhanceSdrToHdr struct {
	value string
}

type QualityEnhanceSdrToHdrEnum struct {
	SD_RTO_HDR10      QualityEnhanceSdrToHdr
	SD_RTO_HDR10_FLAT QualityEnhanceSdrToHdr
}

func GetQualityEnhanceSdrToHdrEnum() QualityEnhanceSdrToHdrEnum {
	return QualityEnhanceSdrToHdrEnum{
		SD_RTO_HDR10: QualityEnhanceSdrToHdr{
			value: "SDRtoHDR10",
		},
		SD_RTO_HDR10_FLAT: QualityEnhanceSdrToHdr{
			value: "SDRtoHDR10FLAT",
		},
	}
}

func (c QualityEnhanceSdrToHdr) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QualityEnhanceSdrToHdr) UnmarshalJSON(b []byte) error {
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
