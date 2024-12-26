package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LivePlayingScriptInfo 直播进度信息。
type LivePlayingScriptInfo struct {

	// 剧本名称
	ScriptName *string `json:"script_name,omitempty"`

	// 数字人模型资产ID，可以从资产库中查询。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]LivePlayingShootScriptItem `json:"shoot_scripts,omitempty"`
}

func (o LivePlayingScriptInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LivePlayingScriptInfo struct{}"
	}

	return strings.Join([]string{"LivePlayingScriptInfo", string(data)}, " ")
}
