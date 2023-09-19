package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateVideoScriptsReq 剧本信息更新。
type UpdateVideoScriptsReq struct {

	// 剧本名称。
	ScriptName *string `json:"script_name,omitempty"`

	// 剧本描述。
	ScriptDescription *string `json:"script_description,omitempty"`

	// 视频生成类型。该参数取值是MODEL时，model_asset_id必填；取值是PICTURE时，human_image必填。 * MODEL：通过分数数字人模型生成视频 * PICTURE： 通过单张照片生成视频
	VideoMakingType *UpdateVideoScriptsReqVideoMakingType `json:"video_making_type,omitempty"`

	// 人物照片，需要Base64编码。
	HumanImage *string `json:"human_image,omitempty"`

	// 数字人资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 数字人模型类型。  * HUMAN_MODEL_2D：分身数字人 * HUMAN_MODEL_3D：3D数字人
	ModelAssetType *UpdateVideoScriptsReqModelAssetType `json:"model_asset_type,omitempty"`

	// 场景资产ID。
	SceneAssetId *string `json:"scene_asset_id,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 剧本参数。
	ShootScripts *[]ShootScriptItem `json:"shoot_scripts,omitempty"`

	// 私有数据，用户填写，原样带回。
	PrivData *string `json:"priv_data,omitempty"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`
}

func (o UpdateVideoScriptsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVideoScriptsReq struct{}"
	}

	return strings.Join([]string{"UpdateVideoScriptsReq", string(data)}, " ")
}

type UpdateVideoScriptsReqVideoMakingType struct {
	value string
}

type UpdateVideoScriptsReqVideoMakingTypeEnum struct {
	MODEL   UpdateVideoScriptsReqVideoMakingType
	PICTURE UpdateVideoScriptsReqVideoMakingType
}

func GetUpdateVideoScriptsReqVideoMakingTypeEnum() UpdateVideoScriptsReqVideoMakingTypeEnum {
	return UpdateVideoScriptsReqVideoMakingTypeEnum{
		MODEL: UpdateVideoScriptsReqVideoMakingType{
			value: "MODEL",
		},
		PICTURE: UpdateVideoScriptsReqVideoMakingType{
			value: "PICTURE",
		},
	}
}

func (c UpdateVideoScriptsReqVideoMakingType) Value() string {
	return c.value
}

func (c UpdateVideoScriptsReqVideoMakingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVideoScriptsReqVideoMakingType) UnmarshalJSON(b []byte) error {
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
