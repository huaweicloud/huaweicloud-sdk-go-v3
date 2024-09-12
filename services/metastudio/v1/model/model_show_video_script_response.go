package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowVideoScriptResponse Response Object
type ShowVideoScriptResponse struct {

	// **参数解释**： 剧本名称。 **约束限制**： 不涉及。 **取值范围**： 只能使用中英文字符，字符长度1-256位。 **默认取值**： 不涉及。
	ScriptName *string `json:"script_name,omitempty"`

	// **参数解释**： 剧本描述。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-1024位。 **默认取值**： 不涉及。
	ScriptDescription *string `json:"script_description,omitempty"`

	// **参数解释**： 横竖屏类型。 **约束限制**： 不涉及。 **取值范围**： * LANDSCAPE：横屏。 * VERTICAL：竖屏。
	ViewMode *ShowVideoScriptResponseViewMode `json:"view_mode,omitempty"`

	// **参数解释**： 数字人模型资产ID。 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// **参数解释**： 数字人模型类型。 **约束限制**： 不涉及 **取值范围**： * HUMAN_MODEL_2D：分身数字人 * HUMAN_MODEL_3D：3D数字人  **默认取值**： 不涉及
	ModelAssetType *ShowVideoScriptResponseModelAssetType `json:"model_asset_type,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// **参数解释**： 场景资产ID。 **约束限制**： 分身数字人视频制作不需要填写该参数。 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	SceneAssetId *string `json:"scene_asset_id,omitempty"`

	// **参数解释**： 私有数据，用户填写，原样带回。 **约束限制**： 不涉及 **取值范围**： 字符长度0-8192位 **默认取值**： 不涉及
	PrivData *string `json:"priv_data,omitempty"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]ShootScriptShowItem `json:"shoot_scripts,omitempty"`

	// 剧本ID。
	ScriptId *string `json:"script_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	AudioFiles *ShootScriptAudioFiles `json:"audio_files,omitempty"`

	// 剧本封面下载url。
	ScriptCoverUrl *string `json:"script_cover_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVideoScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoScriptResponse struct{}"
	}

	return strings.Join([]string{"ShowVideoScriptResponse", string(data)}, " ")
}

type ShowVideoScriptResponseViewMode struct {
	value string
}

type ShowVideoScriptResponseViewModeEnum struct {
	LANDSCAPE ShowVideoScriptResponseViewMode
	VERTICAL  ShowVideoScriptResponseViewMode
}

func GetShowVideoScriptResponseViewModeEnum() ShowVideoScriptResponseViewModeEnum {
	return ShowVideoScriptResponseViewModeEnum{
		LANDSCAPE: ShowVideoScriptResponseViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: ShowVideoScriptResponseViewMode{
			value: "VERTICAL",
		},
	}
}

func (c ShowVideoScriptResponseViewMode) Value() string {
	return c.value
}

func (c ShowVideoScriptResponseViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVideoScriptResponseViewMode) UnmarshalJSON(b []byte) error {
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

type ShowVideoScriptResponseModelAssetType struct {
	value string
}

type ShowVideoScriptResponseModelAssetTypeEnum struct {
	HUMAN_MODEL_2_D ShowVideoScriptResponseModelAssetType
	HUMAN_MODEL_3_D ShowVideoScriptResponseModelAssetType
}

func GetShowVideoScriptResponseModelAssetTypeEnum() ShowVideoScriptResponseModelAssetTypeEnum {
	return ShowVideoScriptResponseModelAssetTypeEnum{
		HUMAN_MODEL_2_D: ShowVideoScriptResponseModelAssetType{
			value: "HUMAN_MODEL_2D",
		},
		HUMAN_MODEL_3_D: ShowVideoScriptResponseModelAssetType{
			value: "HUMAN_MODEL_3D",
		},
	}
}

func (c ShowVideoScriptResponseModelAssetType) Value() string {
	return c.value
}

func (c ShowVideoScriptResponseModelAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVideoScriptResponseModelAssetType) UnmarshalJSON(b []byte) error {
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
