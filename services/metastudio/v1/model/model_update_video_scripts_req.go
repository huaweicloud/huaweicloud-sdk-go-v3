package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateVideoScriptsReq 剧本信息更新。
type UpdateVideoScriptsReq struct {

	// **参数解释**： 剧本名称。 **约束限制**： 不涉及。 **取值范围**： 只能使用中英文字符，字符长度1-256位。 **默认取值**： 不涉及。
	ScriptName *string `json:"script_name,omitempty"`

	// **参数解释**： 剧本描述。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-1024位。 **默认取值**： 不涉及。
	ScriptDescription *string `json:"script_description,omitempty"`

	// **参数解释**： 横竖屏类型。 **约束限制**： 不涉及。 **取值范围**： * LANDSCAPE：横屏。 * VERTICAL：竖屏。
	ViewMode *UpdateVideoScriptsReqViewMode `json:"view_mode,omitempty"`

	// **参数解释**： 数字人模型资产ID。 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// **参数解释**： 数字人模型类型。 **约束限制**： 不涉及 **取值范围**： * HUMAN_MODEL_2D：分身数字人 * HUMAN_MODEL_3D：3D数字人  **默认取值**： 不涉及
	ModelAssetType *UpdateVideoScriptsReqModelAssetType `json:"model_asset_type,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// **参数解释**： 场景资产ID。 **约束限制**： 分身数字人视频制作不需要填写该参数。 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	SceneAssetId *string `json:"scene_asset_id,omitempty"`

	// **参数解释**： 私有数据，用户填写，原样带回。 **约束限制**： 不涉及 **取值范围**： 字符长度0-8192位 **默认取值**： 不涉及
	PrivData *string `json:"priv_data,omitempty"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]ShootScriptItem `json:"shoot_scripts,omitempty"`
}

func (o UpdateVideoScriptsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVideoScriptsReq struct{}"
	}

	return strings.Join([]string{"UpdateVideoScriptsReq", string(data)}, " ")
}

type UpdateVideoScriptsReqViewMode struct {
	value string
}

type UpdateVideoScriptsReqViewModeEnum struct {
	LANDSCAPE UpdateVideoScriptsReqViewMode
	VERTICAL  UpdateVideoScriptsReqViewMode
}

func GetUpdateVideoScriptsReqViewModeEnum() UpdateVideoScriptsReqViewModeEnum {
	return UpdateVideoScriptsReqViewModeEnum{
		LANDSCAPE: UpdateVideoScriptsReqViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: UpdateVideoScriptsReqViewMode{
			value: "VERTICAL",
		},
	}
}

func (c UpdateVideoScriptsReqViewMode) Value() string {
	return c.value
}

func (c UpdateVideoScriptsReqViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVideoScriptsReqViewMode) UnmarshalJSON(b []byte) error {
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

type UpdateVideoScriptsReqModelAssetType struct {
	value string
}

type UpdateVideoScriptsReqModelAssetTypeEnum struct {
	HUMAN_MODEL_2_D UpdateVideoScriptsReqModelAssetType
	HUMAN_MODEL_3_D UpdateVideoScriptsReqModelAssetType
}

func GetUpdateVideoScriptsReqModelAssetTypeEnum() UpdateVideoScriptsReqModelAssetTypeEnum {
	return UpdateVideoScriptsReqModelAssetTypeEnum{
		HUMAN_MODEL_2_D: UpdateVideoScriptsReqModelAssetType{
			value: "HUMAN_MODEL_2D",
		},
		HUMAN_MODEL_3_D: UpdateVideoScriptsReqModelAssetType{
			value: "HUMAN_MODEL_3D",
		},
	}
}

func (c UpdateVideoScriptsReqModelAssetType) Value() string {
	return c.value
}

func (c UpdateVideoScriptsReqModelAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVideoScriptsReqModelAssetType) UnmarshalJSON(b []byte) error {
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
