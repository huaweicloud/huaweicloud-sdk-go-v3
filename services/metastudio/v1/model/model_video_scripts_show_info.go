package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoScriptsShowInfo struct {

	// 剧本名称
	ScriptName *string `json:"script_name,omitempty"`

	// 剧本描述。
	ScriptDescription *string `json:"script_description,omitempty"`

	// 横竖屏类型。默认值是LANDSCAPE。 * LANDSCAPE：横屏。 * VERTICAL： 竖屏。
	ViewMode *VideoScriptsShowInfoViewMode `json:"view_mode,omitempty"`

	// 数字人模型资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 数字人模型类型。  * HUMAN_MODEL_2D：分身数字人 * HUMAN_MODEL_3D：3D数字人
	ModelAssetType *VideoScriptsShowInfoModelAssetType `json:"model_asset_type,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 场景资产ID。 > * 分身数字人视频制作不需要填写该参数。
	SceneAssetId *string `json:"scene_asset_id,omitempty"`

	// 私有数据，用户填写，原样带回。
	PrivData *string `json:"priv_data,omitempty"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]ShootScriptShowItem `json:"shoot_scripts,omitempty"`
}

func (o VideoScriptsShowInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoScriptsShowInfo struct{}"
	}

	return strings.Join([]string{"VideoScriptsShowInfo", string(data)}, " ")
}

type VideoScriptsShowInfoViewMode struct {
	value string
}

type VideoScriptsShowInfoViewModeEnum struct {
	LANDSCAPE VideoScriptsShowInfoViewMode
	VERTICAL  VideoScriptsShowInfoViewMode
}

func GetVideoScriptsShowInfoViewModeEnum() VideoScriptsShowInfoViewModeEnum {
	return VideoScriptsShowInfoViewModeEnum{
		LANDSCAPE: VideoScriptsShowInfoViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: VideoScriptsShowInfoViewMode{
			value: "VERTICAL",
		},
	}
}

func (c VideoScriptsShowInfoViewMode) Value() string {
	return c.value
}

func (c VideoScriptsShowInfoViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoScriptsShowInfoViewMode) UnmarshalJSON(b []byte) error {
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

type VideoScriptsShowInfoModelAssetType struct {
	value string
}

type VideoScriptsShowInfoModelAssetTypeEnum struct {
	HUMAN_MODEL_2_D VideoScriptsShowInfoModelAssetType
	HUMAN_MODEL_3_D VideoScriptsShowInfoModelAssetType
}

func GetVideoScriptsShowInfoModelAssetTypeEnum() VideoScriptsShowInfoModelAssetTypeEnum {
	return VideoScriptsShowInfoModelAssetTypeEnum{
		HUMAN_MODEL_2_D: VideoScriptsShowInfoModelAssetType{
			value: "HUMAN_MODEL_2D",
		},
		HUMAN_MODEL_3_D: VideoScriptsShowInfoModelAssetType{
			value: "HUMAN_MODEL_3D",
		},
	}
}

func (c VideoScriptsShowInfoModelAssetType) Value() string {
	return c.value
}

func (c VideoScriptsShowInfoModelAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoScriptsShowInfoModelAssetType) UnmarshalJSON(b []byte) error {
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
