package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTtsaReq struct {

	// 音色模型ID
	VoiceAssetId string `json:"voice_asset_id"`

	// HTML格式的台词，可包含动作。最多2048个字符。 > * HTML格式举例：\\<speak>大家好<insert-action id=\\\"14cc7bbcde4982aab82f9d9af9e0f743\\\"/>，非常高兴给大家介绍MetaStudio。\\</speak> > * insert-action id通过查询资产列表接口获取，查询时asset_type=ANIMATION > * 多音字标签：\\<phoneme ph=\\\"拼音\\\">汉字\\</phoneme>，南京\\<phoneme ph=\\\"shi4 zhang3\\\">市长\\</phoneme>江大桥。 > * 停顿标签：\\<break/>，中方一贯主张\\<break/>维护国家主权平等，不干涉他国内政\\<break time=\\\"300ms\\\"/>是联合国宪章\\<break time=\\\"500ms\\\"/>最重要的原则。
	Text string `json:"text"`

	// 语速。  取值范围[50,200]   默认值：100
	Speed *int32 `json:"speed,omitempty"`

	// 基频。  取值范围[50,200]  默认值：100
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。  取值范围[90,240]   默认值：100
	Volume *int32 `json:"volume,omitempty"`

	// 情感标签。 * ANGER：愤怒 * HAPPY：开心 * SAD：悲伤 * CALM：平静
	Emotion *string `json:"emotion,omitempty"`

	// 关联父任务任务ID。
	ParentJobId *string `json:"parent_job_id,omitempty"`

	// 是否生成智能动作数据。
	AutoMotion *bool `json:"auto_motion,omitempty"`

	// 风格化ID。
	StyleId *string `json:"style_id,omitempty"`

	// 人位置及相机位置。由如下4组浮点数组成的字符：人位置的X/Y/Z值，人角度的Pitch/Yaw/Roll值；相机位置的X/Y/Z值，相机角度的Pitch/Yaw/Roll值。
	CameraPosition *string `json:"camera_position,omitempty"`
}

func (o CreateTtsaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsaReq struct{}"
	}

	return strings.Join([]string{"CreateTtsaReq", string(data)}, " ")
}
