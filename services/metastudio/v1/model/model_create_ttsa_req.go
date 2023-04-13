package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTtsaReq struct {

	// 音色模型资产ID。
	VoiceAssetId string `json:"voice_asset_id"`

	// HTML格式的台词，可包含动作。最多2048个字符。 > * HTML格式举例：\\<speak>大家好<insert-action id=\\\"14cc7bbcde4982aab82f9d9af9e0f743\\\"/>，非常高兴给大家介绍MetaStudio。\\</speak> > * insert-action id通过[查询资产列表]接口获取，查询时asset_type=ANIMATION
	Text string `json:"text"`

	// 语速。最小值50，最大值200，默认值100。
	Speed *int32 `json:"speed,omitempty"`

	// 基频。最小值50，最大值200，默认值100。
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。最小值90，最大值240，默认值100。
	Volume *int32 `json:"volume,omitempty"`

	// 情感标签。 * ANGER: 愤怒 * HAPPY: 开心 * SAD: 悲伤 * CALM: 平静
	Emotion *string `json:"emotion,omitempty"`

	// 关联父任务任务ID。
	ParentJobId *string `json:"parent_job_id,omitempty"`

	// 是否生成智能动作数据
	AutoMotion *bool `json:"auto_motion,omitempty"`

	// 风格化id
	StyleId *string `json:"style_id,omitempty"`
}

func (o CreateTtsaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsaReq struct{}"
	}

	return strings.Join([]string{"CreateTtsaReq", string(data)}, " ")
}
