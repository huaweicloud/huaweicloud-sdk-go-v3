package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePhotoDigitalHumanVideoReq struct {

	// 剧本ID。 > * 如果shoot_scripts中shoot_script.script_type为\"TEXT\"，则台词以shoot_scripts中的文本为准； > * 如果shoot_scripts中shoot_script.script_type为\"AUDIO\"，则台词以script_id对应剧本中的音频为准。
	ScriptId *string `json:"script_id,omitempty"`

	// 人物照片，需要Base64编码。照片分辨率不超过1080P。
	HumanImage string `json:"human_image"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *PhotoVideoConfig `json:"video_config,omitempty"`

	// 剧本列表。照片数字人仅支持传入一个剧本shoot_script，剧本参数仅支持shoot_script.script_type、shoot_script.text_config；
	ShootScripts []ShootScriptItem `json:"shoot_scripts"`

	OutputAssetConfig *OutputAssetConfig `json:"output_asset_config"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	CallbackConfig *CallBackConfig `json:"callback_config,omitempty"`
}

func (o CreatePhotoDigitalHumanVideoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePhotoDigitalHumanVideoReq struct{}"
	}

	return strings.Join([]string{"CreatePhotoDigitalHumanVideoReq", string(data)}, " ")
}
