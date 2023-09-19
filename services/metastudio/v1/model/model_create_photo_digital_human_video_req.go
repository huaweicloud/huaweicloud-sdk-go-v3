package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePhotoDigitalHumanVideoReq struct {

	// 剧本ID。 > * 如果填写了script_id，model_asset_id、voice_config、scene_asset_id、video_config、shoot_scripts可以不填，以脚本中的配置为准。 > * 如果填写了script_id，并且同时也填写了model_asset_id、voice_config、scene_asset_id、video_config、shoot_scripts则以本接口中的配置为准。
	ScriptId *string `json:"script_id,omitempty"`

	// 人物照片，需要Base64编码。照片分辨率不超过1080P。
	HumanImage string `json:"human_image"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *PhotoVideoConfig `json:"video_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts []ShootScriptItem `json:"shoot_scripts"`

	OutputAssetConfig *OutputAssetConfig `json:"output_asset_config"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`
}

func (o CreatePhotoDigitalHumanVideoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePhotoDigitalHumanVideoReq struct{}"
	}

	return strings.Join([]string{"CreatePhotoDigitalHumanVideoReq", string(data)}, " ")
}
