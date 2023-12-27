package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateVideoScriptsReq 剧本信息更新。
type UpdateVideoScriptsReq struct {

	// 剧本名称
	ScriptName *string `json:"script_name,omitempty"`

	// 剧本描述。
	ScriptDescription *string `json:"script_description,omitempty"`

	// 数字人模型资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 数字人模型类型。  * HUMAN_MODEL_2D：分身数字人 * HUMAN_MODEL_3D：3D数字人
	ModelAssetType *UpdateVideoScriptsReqModelAssetType `json:"model_asset_type,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 场景资产ID。 > * 分身数字人视频制作不需要填写该参数。
	SceneAssetId *string `json:"scene_asset_id,omitempty"`

	// 私有数据，用户填写，原样带回。
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
