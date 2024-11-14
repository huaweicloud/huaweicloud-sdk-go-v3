package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTtsaReq struct {

	// 音色模型ID。需要使用MetaStudio的数字资产管理相关接口从资产库查出。
	VoiceAssetId string `json:"voice_asset_id"`

	// 脚本类型，即视频制作的驱动方式。默认TEXT * TEXT: 文本驱动，即通过TTS合成语音 * AUDIO: 语音驱动
	ScriptType *CreateTtsaReqScriptType `json:"script_type,omitempty"`

	// HTML格式的台词，可包含动作。最多2048个字符。 > * HTML格式举例：\\<speak>大家好<insert-action id=\\\"14cc7bbcde4982aab82f9d9af9e0f743\\\"/>，非常高兴给大家介绍MetaStudio。\\</speak> > * insert-action id通过查询资产列表接口获取，查询时asset_type=ANIMATION > * 多音字标签：\\<phoneme ph=\\\"拼音\\\">汉字\\</phoneme>，南京\\<phoneme ph=\\\"shi4 zhang3\\\">市长\\</phoneme>江大桥。 > * 停顿标签：\\<break/>，中方一贯主张\\<break/>维护国家主权平等，不干涉他国内政\\<break time=\\\"300ms\\\"/>是联合国宪章\\<break time=\\\"500ms\\\"/>最重要的原则。
	Text *string `json:"text,omitempty"`

	// 语音驱动音频文件下载URL。
	AudioFileDownloadUrl *string `json:"audio_file_download_url,omitempty"`

	// 语速。  取值范围[50,200]   默认值：100
	Speed *int32 `json:"speed,omitempty"`

	// 基频。  取值范围[50,200]  默认值：100
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。  取值范围[90,240]   默认值：100
	Volume *int32 `json:"volume,omitempty"`

	// 情感标签。 * ANGER：愤怒 * HAPPY：开心 * SAD：悲伤 * CALM：平静
	Emotion *string `json:"emotion,omitempty"`

	// 风格化ID。需要调用数字人风格管理相关接口，从系统中查得。
	StyleId string `json:"style_id"`

	// 人位置及相机位置。由如下4组浮点数组成的字符：人位置的X/Y/Z值，人角度的Pitch/Yaw/Roll值；相机位置的X/Y/Z值，相机角度的Pitch/Yaw/Roll值。
	CameraPosition *string `json:"camera_position,omitempty"`

	// 任务类型。 * REAL_JOB：实时任务。如数字人交互。 * UNREAL_JOB：非实时任务。如数字人视频制作
	JobType *CreateTtsaReqJobType `json:"job_type,omitempty"`
}

func (o CreateTtsaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsaReq struct{}"
	}

	return strings.Join([]string{"CreateTtsaReq", string(data)}, " ")
}

type CreateTtsaReqScriptType struct {
	value string
}

type CreateTtsaReqScriptTypeEnum struct {
	TEXT  CreateTtsaReqScriptType
	AUDIO CreateTtsaReqScriptType
}

func GetCreateTtsaReqScriptTypeEnum() CreateTtsaReqScriptTypeEnum {
	return CreateTtsaReqScriptTypeEnum{
		TEXT: CreateTtsaReqScriptType{
			value: "TEXT",
		},
		AUDIO: CreateTtsaReqScriptType{
			value: "AUDIO",
		},
	}
}

func (c CreateTtsaReqScriptType) Value() string {
	return c.value
}

func (c CreateTtsaReqScriptType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTtsaReqScriptType) UnmarshalJSON(b []byte) error {
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

type CreateTtsaReqJobType struct {
	value string
}

type CreateTtsaReqJobTypeEnum struct {
	REAL_JOB   CreateTtsaReqJobType
	UNREAL_JOB CreateTtsaReqJobType
}

func GetCreateTtsaReqJobTypeEnum() CreateTtsaReqJobTypeEnum {
	return CreateTtsaReqJobTypeEnum{
		REAL_JOB: CreateTtsaReqJobType{
			value: "REAL_JOB",
		},
		UNREAL_JOB: CreateTtsaReqJobType{
			value: "UNREAL_JOB",
		},
	}
}

func (c CreateTtsaReqJobType) Value() string {
	return c.value
}

func (c CreateTtsaReqJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTtsaReqJobType) UnmarshalJSON(b []byte) error {
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
