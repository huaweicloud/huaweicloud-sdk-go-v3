package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveVideoScriptInfo 创视频制作脚本请求。
type LiveVideoScriptInfo struct {

	// 剧本名称
	ScriptName string `json:"script_name"`

	// 剧本描述。
	ScriptDescription *string `json:"script_description,omitempty"`

	// 数字人ID。对应形象和音色组合。
	DhId *string `json:"dh_id,omitempty"`

	// 数字人模型资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	// 背景配置。
	BackgroundConfig *[]BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]LiveShootScriptItem `json:"shoot_scripts,omitempty"`
}

func (o LiveVideoScriptInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveVideoScriptInfo struct{}"
	}

	return strings.Join([]string{"LiveVideoScriptInfo", string(data)}, " ")
}
