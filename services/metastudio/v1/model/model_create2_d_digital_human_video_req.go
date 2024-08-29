package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Create2DDigitalHumanVideoReq struct {

	// 剧本ID。 > * 如果填写了script_id，model_asset_id、voice_config、scene_asset_id、video_config、shoot_scripts可以不填，以脚本中的配置为准。 > * 如果填写了script_id，并且同时也填写了model_asset_id、voice_config、scene_asset_id、video_config、shoot_scripts则以本接口中的配置为准。
	ScriptId *string `json:"script_id,omitempty"`

	// 分身数字人模型资产ID，可以从资产库中查询。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]ShootScriptItem `json:"shoot_scripts,omitempty"`

	OutputAssetConfig *OutputAssetConfig `json:"output_asset_config,omitempty"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	CallbackConfig *CallBackConfig `json:"callback_config,omitempty"`
}

func (o Create2DDigitalHumanVideoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Create2DDigitalHumanVideoReq struct{}"
	}

	return strings.Join([]string{"Create2DDigitalHumanVideoReq", string(data)}, " ")
}
