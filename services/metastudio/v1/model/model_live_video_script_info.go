package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveVideoScriptInfo 创视频制作脚本请求。
type LiveVideoScriptInfo struct {

	// **参数解释**： 剧本ID。 **约束限制**： 该字段无需填写。 **取值范围**： 字符长度1-64位。 **默认取值**： 不涉及。
	ScriptId *string `json:"script_id,omitempty"`

	// **参数解释**： 剧本名称。 **约束限制**： 该字段必须填写。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及。
	ScriptName string `json:"script_name"`

	// **参数解释**： 剧本描述。 **约束限制**： 该字段无需填写。 **取值范围**： 字符长度0-1024位。 **默认取值**： 不涉及。
	ScriptDescription *string `json:"script_description,omitempty"`

	// **参数解释**： 数字人ID。对应形象和音色组合。 **约束限制**： 该字段暂未启用，无需填写。 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及。
	DhId *string `json:"dh_id,omitempty"`

	// **参数解释**： 数字人模型资产ID，可以从资产库中查询。 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	// 背景配置。
	BackgroundConfig *[]BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts []LiveShootScriptItem `json:"shoot_scripts"`
}

func (o LiveVideoScriptInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveVideoScriptInfo struct{}"
	}

	return strings.Join([]string{"LiveVideoScriptInfo", string(data)}, " ")
}
